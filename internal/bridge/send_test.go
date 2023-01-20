// Copyright (c) 2023 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package bridge_test

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/ProtonMail/go-proton-api"
	"github.com/ProtonMail/go-proton-api/server"
	"github.com/ProtonMail/proton-bridge/v3/internal/bridge"
	"github.com/ProtonMail/proton-bridge/v3/internal/constants"
	"github.com/ProtonMail/proton-bridge/v3/internal/events"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/stretchr/testify/require"
)

func TestBridge_Send(t *testing.T) {
	withEnv(t, func(ctx context.Context, s *server.Server, netCtl *proton.NetCtl, locator bridge.Locator, storeKey []byte) {
		_, _, err := s.CreateUser("recipient", password)
		require.NoError(t, err)

		withBridge(ctx, t, s.GetHostURL(), netCtl, locator, storeKey, func(bridge *bridge.Bridge, _ *bridge.Mocks) {
			senderUserID, err := bridge.LoginFull(ctx, username, password, nil, nil)
			require.NoError(t, err)

			recipientUserID, err := bridge.LoginFull(ctx, "recipient", password, nil, nil)
			require.NoError(t, err)

			senderInfo, err := bridge.GetUserInfo(senderUserID)
			require.NoError(t, err)

			recipientInfo, err := bridge.GetUserInfo(recipientUserID)
			require.NoError(t, err)

			for i := 0; i < 10; i++ {
				// Dial the server.
				client, err := smtp.Dial(net.JoinHostPort(constants.Host, fmt.Sprint(bridge.GetSMTPPort())))
				require.NoError(t, err)
				defer client.Close() //nolint:errcheck

				// Upgrade to TLS.
				require.NoError(t, client.StartTLS(&tls.Config{InsecureSkipVerify: true}))

				if i%2 == 0 {
					// Authorize with SASL PLAIN.
					require.NoError(t, client.Auth(sasl.NewPlainClient(
						senderInfo.Addresses[0],
						senderInfo.Addresses[0],
						string(senderInfo.BridgePass)),
					))
				} else {
					// Authorize with SASL LOGIN.
					require.NoError(t, client.Auth(sasl.NewLoginClient(
						senderInfo.Addresses[0],
						string(senderInfo.BridgePass)),
					))
				}

				// Send the message.
				require.NoError(t, client.SendMail(
					senderInfo.Addresses[0],
					[]string{recipientInfo.Addresses[0]},
					strings.NewReader(fmt.Sprintf("Subject: Test %v\r\n\r\nHello world!", i)),
				))
			}

			// Connect the sender IMAP client.
			senderIMAPClient, err := client.Dial(net.JoinHostPort(constants.Host, fmt.Sprint(bridge.GetIMAPPort())))
			require.NoError(t, err)
			require.NoError(t, senderIMAPClient.Login(senderInfo.Addresses[0], string(senderInfo.BridgePass)))
			defer senderIMAPClient.Logout() //nolint:errcheck

			// Connect the recipient IMAP client.
			recipientIMAPClient, err := client.Dial(net.JoinHostPort(constants.Host, fmt.Sprint(bridge.GetIMAPPort())))
			require.NoError(t, err)
			require.NoError(t, recipientIMAPClient.Login(recipientInfo.Addresses[0], string(recipientInfo.BridgePass)))
			defer recipientIMAPClient.Logout() //nolint:errcheck

			// Sender should have 10 messages in the sent folder.
			// Recipient should have 0 messages in inbox.
			require.Eventually(t, func() bool {
				sent, err := senderIMAPClient.Status(`Sent`, []imap.StatusItem{imap.StatusMessages})
				require.NoError(t, err)

				inbox, err := recipientIMAPClient.Status(`Inbox`, []imap.StatusItem{imap.StatusMessages})
				require.NoError(t, err)

				return sent.Messages == 10 && inbox.Messages == 10
			}, 10*time.Second, 100*time.Millisecond)
		})
	})
}

func TestBridge_SendDraftFlags(t *testing.T) {
	withEnv(t, func(ctx context.Context, s *server.Server, netCtl *proton.NetCtl, locator bridge.Locator, storeKey []byte) {
		// Create a recipient user.
		_, _, err := s.CreateUser("recipient", password)
		require.NoError(t, err)

		// The sender should be fully synced.
		withBridge(ctx, t, s.GetHostURL(), netCtl, locator, storeKey, func(bridge *bridge.Bridge, _ *bridge.Mocks) {
			syncCh, done := chToType[events.Event, events.SyncFinished](bridge.GetEvents(events.SyncFinished{}))
			defer done()

			userID, err := bridge.LoginFull(ctx, username, password, nil, nil)
			require.NoError(t, err)

			require.Equal(t, userID, (<-syncCh).UserID)
		})

		// Start the bridge.
		withBridge(ctx, t, s.GetHostURL(), netCtl, locator, storeKey, func(bridge *bridge.Bridge, _ *bridge.Mocks) {
			// Get the sender user info.
			userInfo, err := bridge.QueryUserInfo(username)
			require.NoError(t, err)

			// Connect the sender IMAP client.
			imapClient, err := client.Dial(net.JoinHostPort(constants.Host, fmt.Sprint(bridge.GetIMAPPort())))
			require.NoError(t, err)
			require.NoError(t, imapClient.Login(userInfo.Addresses[0], string(userInfo.BridgePass)))
			defer imapClient.Logout() //nolint:errcheck

			// The message to send.
			const message = `Subject: Test\r\n\r\nHello world!`

			// Save a draft.
			require.NoError(t, imapClient.Append("Drafts", []string{imap.DraftFlag}, time.Now(), strings.NewReader(message)))

			// Assert that the draft exists and is marked as a draft.
			{
				messages, err := clientFetch(imapClient, "Drafts")
				require.NoError(t, err)
				require.Len(t, messages, 1)
				require.Contains(t, messages[0].Flags, imap.DraftFlag)
			}

			// Connect the SMTP client.
			smtpClient, err := smtp.Dial(net.JoinHostPort(constants.Host, fmt.Sprint(bridge.GetSMTPPort())))
			require.NoError(t, err)
			defer smtpClient.Close() //nolint:errcheck

			// Upgrade to TLS.
			require.NoError(t, smtpClient.StartTLS(&tls.Config{InsecureSkipVerify: true}))

			// Authorize with SASL PLAIN.
			require.NoError(t, smtpClient.Auth(sasl.NewPlainClient(
				userInfo.Addresses[0],
				userInfo.Addresses[0],
				string(userInfo.BridgePass)),
			))

			// Send the message.
			require.NoError(t, smtpClient.SendMail(
				userInfo.Addresses[0],
				[]string{"recipient@" + s.GetDomain()},
				strings.NewReader(message),
			))

			// Delete the draft: add the \Deleted flag and expunge.
			{
				status, err := imapClient.Select("Drafts", false)
				require.NoError(t, err)
				require.Equal(t, uint32(1), status.Messages)

				// Add the \Deleted flag.
				require.NoError(t, clientStore(imapClient, 1, 1, true, imap.FormatFlagsOp(imap.AddFlags, true), imap.DeletedFlag))

				// Expunge.
				require.NoError(t, imapClient.Expunge(nil))
			}

			// Assert that the draft is eventually gone.
			require.Eventually(t, func() bool {
				status, err := imapClient.Select("Drafts", false)
				require.NoError(t, err)
				return status.Messages == 0
			}, 10*time.Second, 100*time.Millisecond)

			// Assert that the message is eventually in the sent folder.
			require.Eventually(t, func() bool {
				messages, err := clientFetch(imapClient, "Sent")
				require.NoError(t, err)
				return len(messages) == 1
			}, 10*time.Second, 100*time.Millisecond)

			// Assert that the message is not marked as a draft.
			{
				messages, err := clientFetch(imapClient, "Sent")
				require.NoError(t, err)
				require.Len(t, messages, 1)
				require.NotContains(t, messages[0].Flags, imap.DraftFlag)
			}
		})
	})
}
