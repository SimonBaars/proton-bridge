// Copyright (c) 2022 Proton AG
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
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.


#include "Pch.h"
#include "GRPCClient.h"
#include "GRPCUtils.h"
#include "QMLBackend.h"
#include "Exception.h"


using namespace google::protobuf;
using namespace grpc;


namespace
{


/// \todo GODT-1673 Decide how to generate/store/share this certificate.

std::string const cert = R"(-----BEGIN CERTIFICATE-----
MIIC5TCCAc2gAwIBAgIJAMUQK0VGexMsMA0GCSqGSIb3DQEBCwUAMBQxEjAQBgNV
BAMMCWxvY2FsaG9zdDAeFw0yMjA2MTQxNjUyNTVaFw0yMjA3MTQxNjUyNTVaMBQx
EjAQBgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC
ggEBAL6T1JQ0jptq512PBLASpCLFB0px7KIzEml0oMUCkVgUF+2cayrvdBXJZnaO
SG+/JPnHDcQ/ecgqkh2Ii6a2x2kWA5KqWiV+bSHp0drXyUGJfM85muLsnrhYwJ83
HHtweoUVebRZvHn66KjaH8nBJ+YVWyYbSUhJezcg6nBSEtkW+I/XUHu4S2C7FUc5
DXPO3yWWZuZ22OZz70DY3uYE/9COuilotuKdj7XgeKDyKIvRXjPFyqGxwnnp6bXC
vWvrQdcxy0wM+vZxew3QtA/Ag9uKJU9owP6noauXw95l49lEVIA5KXVNtdaldVht
MO/QoelLZC7h79PK22zbii3x930CAwEAAaM6MDgwFAYDVR0RBA0wC4IJbG9jYWxo
b3N0MAsGA1UdDwQEAwIHgDATBgNVHSUEDDAKBggrBgEFBQcDATANBgkqhkiG9w0B
AQsFAAOCAQEAW/9PE8dcAN+0C3K96Xd6Y3qOOtQhRw+WlZXhtiqMtlJfTjvuGKs9
58xuKcTvU5oobxLv+i5+4gpqLjUZZ9FBnYXZIACNVzq4PEXf+YdzcA+y6RS/rqT4
dUjsuYrScAmdXK03Duw3HWYrTp8gsJzIaYGTltUrOn0E4k/TsZb/tZ6z+oH7Fi+p
wdsI6Ut6Zwm3Z7WLn5DDk8KvFjHjZkdsCb82SFSAUVrzWo5EtbLIY/7y3A5rGp9D
t0AVpuGPo5Vn+MW1WA9HT8lhjz0v5wKGMOBi3VYW+Yx8FWHDpacvbZwVM0MjMSAd
M7SXYbNDiLF4LwPLsunoLsW133Ky7s99MA==
-----END CERTIFICATE-----)";


Empty empty; // re-used across client calls.


}


//****************************************************************************************************************************************************
/// \param[out] outError If the function returns false, this variable contains a description of the error.
/// \return true iff the connection was successful.
//****************************************************************************************************************************************************
bool GRPCClient::connectToServer(QString &outError)
{
    try
    {
        SslCredentialsOptions opts;
        opts.pem_root_certs += cert;

        channel_ = CreateChannel("localhost:9292", grpc::SslCredentials(opts));
        if (!channel_)
            throw Exception("Channel creation failed.");

        stub_ = Bridge::NewStub(channel_);
        if (!stub_)
            throw Exception("Stub creation failed.");

        if (!channel_->WaitForConnected(gpr_time_add(gpr_now(GPR_CLOCK_REALTIME),
            gpr_time_from_seconds(10, GPR_TIMESPAN))))
            throw Exception("Connection to the RPC server failed.");

        if (channel_->GetState(true) != GRPC_CHANNEL_READY)
            throw Exception("connection check failed.");

        QMLBackend *backend = &app().backend();
        QObject::connect(this, &GRPCClient::loginFreeUserError, backend, &QMLBackend::loginFreeUserError);
        return true;
    }
    catch (Exception const &e)
    {
        outError = e.qwhat();
        return false;
    }
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::guiReady()
{
    grpc::ClientContext ctx;
    return stub_->GuiReady(&ctx, empty, &empty);
}


//****************************************************************************************************************************************************
/// \param[out] outIsFirst The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isFirstGUIStart(bool &outIsFirst)
{
    return this->getBool(&Bridge::Stub::IsFirstGuiStart, outIsFirst);
}


//****************************************************************************************************************************************************
/// \param[out] outIsOn The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isAutostartOn(bool &outIsOn)
{
    return this->getBool(&Bridge::Stub::IsAutostartOn, outIsOn);
}


//****************************************************************************************************************************************************
/// \param[in] on The new value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setIsAutostartOn(bool on)
{
    return this->setBool(&Bridge::Stub::SetIsAutostartOn, on);
}


//****************************************************************************************************************************************************
/// \param[out] outEnabled The new value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isBetaEnabled(bool &outEnabled)
{
    return this->getBool(&Bridge::Stub::IsBetaEnabled, outEnabled);
}


//****************************************************************************************************************************************************
/// \param[in] enabled The new value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setisBetaEnabled(bool enabled)
{
    return this->setBool(&Bridge::Stub::SetIsBetaEnabled, enabled);
}


//****************************************************************************************************************************************************
/// \param[out] outName The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::colorSchemeName(QString &outName)
{
    return this->getString(&Bridge::Stub::ColorSchemeName, outName);
}


//****************************************************************************************************************************************************
/// \param[in] name The new value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setColorSchemeName(QString const &name)
{
    return this->setString(&Bridge::Stub::SetColorSchemeName, name);
}


//****************************************************************************************************************************************************
/// \param[out] outName The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::currentEmailClient(QString &outName)
{
    return this->getString(&Bridge::Stub::CurrentEmailClient, outName);
}


//****************************************************************************************************************************************************
/// \param[out] outUseSSL The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::useSSLForSMTP(bool &outUseSSL)
{
    return this->getBool(&Bridge::Stub::UseSslForSmtp, outUseSSL);
}


//****************************************************************************************************************************************************
/// \param[in] useSSL The new value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setUseSSLForSMTP(bool useSSL)
{
    return this->setBool(&Bridge::Stub::SetUseSslForSmtp, useSSL);
}


//****************************************************************************************************************************************************
/// \param[out] outPort The port.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::portIMAP(int &outPort)
{
    return this->getInt32(&Bridge::Stub::ImapPort, outPort);
}


//****************************************************************************************************************************************************
/// \param[out] outPort The port.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::portSMTP(int &outPort)
{
    return this->getInt32(&Bridge::Stub::SmtpPort, outPort);
}


//****************************************************************************************************************************************************
/// \param[in] portIMAP The IMAP port.
/// \param[in] portSMTP The SMTP port.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::changePorts(int portIMAP, int portSMTP)
{
    ClientContext ctx;
    ChangePortsRequest request;
    request.set_imapport(portIMAP);
    request.set_smtpport(portSMTP);
    return stub_->ChangePorts(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] outEnabled The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isDoHEnabled(bool &outEnabled)
{
    return this->getBool(&Bridge::Stub::IsDoHEnabled, outEnabled);
}


//****************************************************************************************************************************************************
/// \param[in] enabled The new value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setIsDoHEnabled(bool enabled)
{
    return this->setBool(&Bridge::Stub::SetIsDoHEnabled, enabled);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::quit()
{
    grpc::ClientContext ctx;
    return stub_->Quit(&ctx, empty, &empty);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::restart()
{
    grpc::ClientContext ctx;
    return stub_->Restart(&ctx, empty, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] port The port to check.
/// \param[out] outFree The result of the check.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isPortFree(qint32 port, bool &outFree)
{
    grpc::ClientContext ctx;
    Int32Value p;
    p.set_value(port);
    BoolValue isFree;
    Status result = stub_->IsPortFree(&ctx, p, &isFree);
    if (result.ok())
        outFree = isFree.value();
    return result;
}


//****************************************************************************************************************************************************
/// \param[out]  outValue The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::showOnStartup(bool &outValue)
{
    return this->getBool(&Bridge::Stub::ShowOnStartup, outValue);
}


//****************************************************************************************************************************************************
/// \param[out] outValue The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::showSplashScreen(bool &outValue)
{
    return this->getBool(&Bridge::Stub::ShowSplashScreen, outValue);
}


//****************************************************************************************************************************************************
/// \param[out] outGoos The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::goos(QString &outGoos)
{
    return this->getString(&Bridge::Stub::GoOs, outGoos);
}


//****************************************************************************************************************************************************
/// \param[out] outPath The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::logsPath(QUrl &outPath)
{
    return this->getURLForLocalFile(&Bridge::Stub::LogsPath, outPath);
}


//****************************************************************************************************************************************************
/// \param[out] outPath The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::licensePath(QUrl &outPath)
{
    return this->getURLForLocalFile(&Bridge::Stub::LicensePath, outPath);
}


//****************************************************************************************************************************************************
/// \param[out] outUrl The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::dependencyLicensesLink(QUrl &outUrl)
{
    return this->getURL(&Bridge::Stub::DependencyLicensesLink, outUrl);
}


//****************************************************************************************************************************************************
/// \param[out] outVersion The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::version(QString &outVersion)
{
    return this->getString(&Bridge::Stub::Version, outVersion);
}


//****************************************************************************************************************************************************
/// \param[out] outHostname The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::hostname(QString &outHostname)
{
    return this->getString(&Bridge::Stub::Hostname, outHostname);
}


//****************************************************************************************************************************************************
/// \param[out] outEnabled The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isCacheOnDiskEnabled(bool &outEnabled)
{
    return getBool(&Bridge::Stub::IsCacheOnDiskEnabled, outEnabled);
}


//****************************************************************************************************************************************************
/// \param[out] outPath The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::diskCachePath(QUrl &outPath)
{
    return this->getURLForLocalFile(&Bridge::Stub::DiskCachePath, outPath);
}


//****************************************************************************************************************************************************
/// \param[in] enabled Should the cache be enabled.
/// \param[in] path The value for the property.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::changeLocalCache(bool enabled, QUrl const &path)
{
    grpc::ClientContext ctx;
    ChangeLocalCacheRequest request;
    request.set_enablediskcache(enabled);
    request.set_diskcachepath(path.path(QUrl::FullyDecoded).toStdString());
    return stub_->ChangeLocalCache(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] username The username.
/// \param[in] password The password.
/// \return the status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::login(QString const &username, QString const &password)
{
    grpc::ClientContext ctx;
    LoginRequest request;
    request.set_username(username.toStdString());
    request.set_password(password.toStdString());
    return stub_->Login(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] username The username.
/// \param[in] code The The 2FA code.
/// \return the status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::login2FA(QString const &username, QString const &code)
{
    grpc::ClientContext ctx;
    LoginRequest request;
    request.set_username(username.toStdString());
    request.set_password(code.toStdString());
    return stub_->Login2FA(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] username The username.
/// \param[in] password The password.
/// \return the status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::login2Passwords(QString const &username, QString const &password)
{
    grpc::ClientContext ctx;
    LoginRequest request;
    request.set_username(username.toStdString());
    request.set_password(password.toStdString());
    return stub_->Login2Passwords(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] username The username.
/// \return the status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::loginAbort(QString const &username)
{
    grpc::ClientContext ctx;
    LoginAbortRequest request;
    request.set_username(username.toStdString());
    return stub_->LoginAbort(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::checkUpdate()
{
    return this->simpleMethod(&Bridge::Stub::CheckUpdate);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::installUpdate()
{
    return this->simpleMethod(&Bridge::Stub::InstallUpdate);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setIsAutomaticUpdateOn(bool on)
{
    return this->setBool(&Bridge::Stub::SetIsAutomaticUpdateOn, on);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::isAutomaticUpdateOn(bool &isOn)
{
    return this->getBool(&Bridge::Stub::IsAutomaticUpdateOn, isOn);
}


//****************************************************************************************************************************************************
/// \param[in] userID The user ID.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::logoutUser(QString const &userID)
{
    return methodWithStringParam(&Bridge::Stub::LogoutUser, userID);
}


//****************************************************************************************************************************************************
/// \param[in] userID The user ID.
/// \return the status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::removeUser(QString const &userID)
{
    return methodWithStringParam(&Bridge::Stub::RemoveUser, userID);
}


//****************************************************************************************************************************************************
/// \param[in] userID The user ID.
/// \param[in] address The email address.
/// \return the status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::configureAppleMail(QString const &userID, QString const &address)
{
    ClientContext ctx;
    ConfigureAppleMailRequest request;
    request.set_userid(userID.toStdString());
    request.set_address(address.toStdString());
    return stub_->ConfigureUserAppleMail(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] userID The userID.
/// \param[in] active the new status for the mode.
/// \return The status for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setUserSplitMode(QString const &userID, bool active)
{
    ClientContext ctx;
    UserSplitModeRequest request;
    request.set_userid(userID.toStdString());
    request.set_active(active);

    return stub_->SetUserSplitMode(&ctx, request, &empty);
}


//****************************************************************************************************************************************************
/// \param[out] outUsers The user list.
/// \return The status code for the gRPC call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getUserList(QList<SPUser> &outUsers)
{
    outUsers.clear();

    ClientContext ctx;
    UserListResponse response;
    Status status = stub_->GetUserList(&ctx, empty, &response);
    if (!status.ok())
        return status;

    for (int i = 0; i < response.users_size(); ++i)
        outUsers.append(parsegrpcUser(response.users(i)));

    return status;
}


//****************************************************************************************************************************************************
/// \param[in] userID The userID.
/// \param[out] outUser The user.
/// \return The status code for the operation.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getUser(QString const &userID, ::SPUser &outUser)
{
    ClientContext ctx;
    StringValue s;
    s.set_value(userID.toStdString());
    grpc::User grpcUser;
    Status status = stub_->GetUser(&ctx, s, &grpcUser);

    if (status.ok())
        outUser = parsegrpcUser(grpcUser);

    return grpc::Status();
}


//****************************************************************************************************************************************************
/// \param[out] outKeychains The list of available keychains.
/// \return The status for the gRPC coll.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::availableKeychains(QStringList &outKeychains)
{
    outKeychains.clear();
    ClientContext ctx;
    AvailableKeychainsResponse response;
    Status status = stub_->AvailableKeychains(&ctx, empty, &response);
    if (!status.ok())
        return status;

    for (int i = 0; i < response.keychains_size(); ++i)
        outKeychains.append(QString::fromStdString(response.keychains(i)));

    return status;
}


//****************************************************************************************************************************************************
/// \param[out] outKeychain The current keychain.
/// \return The status for the gRPC coll.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::currentKeychain(QString &outKeychain)
{
    return this->getString(&Bridge::Stub::CurrentKeychain, outKeychain);
}


//****************************************************************************************************************************************************
/// \param[in] keychain The new current keychain.
/// \return The status for the gRPC coll.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setCurrentKeychain(QString const &keychain)
{
    return this->setString(&Bridge::Stub::SetCurrentKeychain, keychain);
}


//****************************************************************************************************************************************************
/// \return The status for the gRPC coll.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::startEventStream()
{
    grpc::ClientContext ctx;
    std::unique_ptr<grpc::ClientReader<grpc::StreamEvent>> reader(stub_->StartEventStream(&ctx, empty));
    grpc::StreamEvent event;

    while (reader->Read(&event))
    {
        switch (event.event_case())
        {
        case grpc::StreamEvent::kApp:
            this->processAppEvent(event.app());
            break;
        case grpc::StreamEvent::kLogin:
            this->processLoginEvent(event.login());
            break;
        case grpc::StreamEvent::kUpdate:
            this->processUpdateEvent(event.update());
            break;
        case grpc::StreamEvent::kCache:
            this->processCacheEvent(event.cache());
            break;
        case grpc::StreamEvent::kMailSettings:
            this->processMailSettingsEvent(event.mailsettings());
            break;
        case grpc::StreamEvent::kKeychain:
            this->processKeychainEvent(event.keychain());
            break;
        case grpc::StreamEvent::kMail:
            this->processMailEvent(event.mail());
            break;
        case grpc::StreamEvent::kUser:
            this->processUserEvent(event.user());
            break;
        default:
            app().log().debug(QString("Unknown stream event type: %1").arg(event.event_case()));
        }
    }

    return reader->Finish();
}


//****************************************************************************************************************************************************
/// \return The status for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::stopEventStream()
{
    grpc::ClientContext ctx;
    return stub_->StopEventStream(&ctx, empty, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] method the gRPC method to call.
/// \return The status for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::simpleMethod(SimpleMethod method)
{
    ClientContext ctx;
    return ((*stub_).*method)(&ctx, empty, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] setter The setter member function.
/// \param[in] value The bool value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setBool(BoolSetter setter, bool value)
{
    grpc::ClientContext ctx;
    BoolValue v;
    v.set_value(value);
    return ((*stub_).*setter)(&ctx, v, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] getter The getter member function.
/// \param[out] outValue The boolean value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getBool(BoolGetter getter, bool &outValue)
{
    grpc::ClientContext ctx;
    BoolValue v;
    Status result = ((*stub_).*getter)(&ctx, empty, &v);
    if (result.ok())
        outValue = v.value();
    return result;
}


//****************************************************************************************************************************************************
/// \param[in] setter The setter member function.
/// \param[in] value The bool value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setInt32(Int32Setter setter, int value)
{
    grpc::ClientContext ctx;
    Int32Value i;
    i.set_value(value);
    return ((*stub_).*setter)(&ctx, i, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] getter The getter member function.
/// \param[out] outValue The boolean value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getInt32(Int32Getter getter, int &outValue)
{
    grpc::ClientContext ctx;
    Int32Value i;
    Status result = ((*stub_).*getter)(&ctx, empty, &i);
    if (result.ok())
        outValue = i.value();
    return result;
}


//****************************************************************************************************************************************************
/// \param[in] setter The setter member function.
/// \param[in] value The string value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::setString(StringSetter setter, QString const &value)
{
    grpc::ClientContext ctx;
    StringValue s;
    s.set_value(value.toStdString());
    return ((*stub_).*setter)(&ctx, s, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] getter The getter member function.
/// \param[out] outValue The string value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getString(StringGetter getter, QString &outValue)
{
    grpc::ClientContext ctx;
    StringValue v;
    Status result = ((*stub_).*getter)(&ctx, empty, &v);
    if (result.ok())
        outValue = QString::fromStdString(v.value());
    return result;
}


//****************************************************************************************************************************************************
/// \param[in] getter The getter member function.
/// \param[out] outValue The URL value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getURLForLocalFile(StringGetter getter, QUrl &outValue)
{
    QString str;
    grpc::Status status = this->getString(getter, str);
    if (status.ok())
        outValue = QUrl::fromLocalFile(str);
    return status;
}


//****************************************************************************************************************************************************
/// \param[in] getter The getter member function.
/// \param[out] outValue The URL value.
/// \return The status code for the call.
//****************************************************************************************************************************************************
grpc::Status GRPCClient::getURL(StringGetter getter, QUrl &outValue)
{
    QString str;
    grpc::Status status = this->getString(getter, str);
    if (status.ok())
        outValue = QUrl(str);
    return status;
}


//****************************************************************************************************************************************************
/// \param[in] method The method to call.
/// \param[in] str The string parameter.
/// \return The status code for the call
//****************************************************************************************************************************************************
grpc::Status GRPCClient::methodWithStringParam(StringParamMethod method, QString const &str)
{
    ClientContext ctx;
    StringValue s;
    s.set_value(str.toStdString());
    return ((*stub_).*method)(&ctx, s, &empty);
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processAppEvent(AppEvent const &event)
{
    switch (event.event_case())
    {
    case AppEvent::kInternetStatus:
        app().log().debug("App event received: InternetStatus.");
        emit internetStatus(event.internetstatus().connected());
        break;
    case AppEvent::kToggleAutostartFinished:
        app().log().debug("App event received: AutostartFinished.");
        emit toggleAutostartFinished();
        break;
    case AppEvent::kResetFinished:
        app().log().debug("App event received: ResetFinished.");
        emit resetFinished();
        break;
    case AppEvent::kReportBugFinished:
        app().log().debug("App event received: ReportBugFinished.");
        emit reportBugFinished();
        break;
    case AppEvent::kReportBugSuccess:
        app().log().debug("App event received: ReportBugSuccess.");
        emit reportBugSuccess();
        break;
    case AppEvent::kReportBugError:
        app().log().debug("App event received: ReportBugError.");
        emit reportBugError();
        break;
    case AppEvent::kShowMainWindow:
        app().log().debug("App event received: ShowMainWindow.");
        emit showMainWindow();
        break;
    default:
        app().log().error("Unknown App event received.");
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processLoginEvent(LoginEvent const &event)
{
    switch (event.event_case())
    {
    case LoginEvent::kError:
    {
        app().log().debug("Login event received: Error.");
        LoginErrorEvent const &error = event.error();
        switch (error.type())
        {
        case USERNAME_PASSWORD_ERROR:
            emit loginUsernamePasswordError(QString::fromStdString(error.message()));
            break;
        case FREE_USER:
            emit loginFreeUserError();
            break;
        case CONNECTION_ERROR:
            emit loginConnectionError(QString::fromStdString(error.message()));
            break;
        case TFA_ERROR:
            emit login2FAError(QString::fromStdString(error.message()));
            break;
        case TFA_ABORT:
            emit login2FAErrorAbort(QString::fromStdString(error.message()));
            break;
        case TWO_PASSWORDS_ERROR:
            emit login2PasswordError(QString::fromStdString(error.message()));
            break;
        case TWO_PASSWORDS_ABORT:
            emit login2PasswordErrorAbort(QString::fromStdString(error.message()));
            break;
        default:
            app().log().debug("Unknown login error event received.");
            break;
        }
        break;
    }
    case LoginEvent::kTfaRequested:
        app().log().debug("Login event received: TfaRequested.");
        emit login2FARequested(QString::fromStdString(event.tfarequested().username()));
        break;
    case LoginEvent::kTwoPasswordRequested:
        app().log().debug("Login event received: TwoPasswordRequested.");
        emit login2PasswordRequested();
        break;
    case LoginEvent::kFinished:
        app().log().debug("Login event received: Finished.");
        emit loginFinished(QString::fromStdString(event.finished().userid()));
        break;
    case LoginEvent::kAlreadyLoggedIn:
        app().log().debug("Login event received: AlreadyLoggedIn.");
        emit loginAlreadyLoggedIn(QString::fromStdString(event.finished().userid()));
        break;
    default:
        app().log().error("Unknown Login event received.");
        break;
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processUpdateEvent(UpdateEvent const &event)
{
    switch (event.event_case())
    {
    case UpdateEvent::kError:
    {
        app().log().debug("Update event received: Error.");

        UpdateErrorEvent const &errorEvent = event.error();
        switch (errorEvent.type())
        {
        case UPDATE_MANUAL_ERROR:
            emit updateManualError();
            break;
        case UPDATE_FORCE_ERROR:
            emit updateForceError();
            break;
        case UPDATE_SILENT_ERROR:
            emit updateSilentError();
            break;
        default:
            app().log().error("Unknown update error received.");
            break;
        }
        break;
    }
    case UpdateEvent::kManualReady:
        app().log().debug("Update event received: ManualReady.");
        emit updateManualReady(QString::fromStdString(event.manualready().version()));
        break;
    case UpdateEvent::kManualRestartNeeded:
        app().log().debug("Update event received: kManualRestartNeeded.");
        emit updateManualRestartNeeded();
        break;
    case UpdateEvent::kForce:
        app().log().debug("Update event received: kForce.");
        emit updateForce(QString::fromStdString(event.force().version()));
        break;
    case UpdateEvent::kSilentRestartNeeded:
        app().log().debug("Update event received: kSilentRestartNeeded.");
        emit updateSilentRestartNeeded();
        break;
    case UpdateEvent::kIsLatestVersion:
        app().log().debug("Update event received: kIsLatestVersion.");
        emit updateIsLatestVersion();
        break;
    case UpdateEvent::kCheckFinished:
        app().log().debug("Update event received: kCheckFinished.");
        emit checkUpdatesFinished();
        break;
    default:
        app().log().error("Unknown Update event received.");
        break;
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processCacheEvent(CacheEvent const &event)
{
    switch (event.event_case())
    {
    case CacheEvent::kError:
    {
        switch (event.error().type())
        {
        case CACHE_UNAVAILABLE_ERROR:
            emit cacheUnavailable();
            break;
        case CACHE_CANT_MOVE_ERROR:
            emit cacheCantMove();
            break;
        case DISK_FULL:
            emit diskFull();
            break;
        default:
            app().log().error("Unknown cache error event received.");
            break;
        }
        break;
    }

    case CacheEvent::kLocationChangedSuccess:
        app().log().debug("Cache event received: LocationChangedSuccess.");
        emit cacheLocationChangeSuccess();
        break;

    case CacheEvent::kChangeLocalCacheFinished:
        emit cacheLocationChangeSuccess();
        app().log().debug("Cache event received: ChangeLocalCacheFinished.");
        break;

    case CacheEvent::kIsCacheOnDiskEnabledChanged:
        app().log().debug("Cache event received: IsCacheOnDiskEnabledChanged.");
        emit isCacheOnDiskEnabledChanged(event.iscacheondiskenabledchanged().enabled());
        break;

    case CacheEvent::kDiskCachePathChanged:
        app().log().debug("Cache event received: DiskCachePathChanged.");
        emit diskCachePathChanged(QUrl::fromLocalFile(QString::fromStdString(event.diskcachepathchanged().path())));
        break;

    default:
        app().log().error("Unknown Cache event received.");
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processMailSettingsEvent(MailSettingsEvent const &event)
{
    switch (event.event_case())
    {
    case MailSettingsEvent::kError:
        app().log().debug("MailSettings event received: Error.");
        switch (event.error().type())
        {
        case IMAP_PORT_ISSUE:
            emit portIssueIMAP();
            break;
        case SMTP_PORT_ISSUE:
            emit portIssueSMTP();
            break;
        default:
            app().log().error("Unknown mail settings error event received.");
            break;
        }

    case MailSettingsEvent::kUseSslForSmtpFinished:
        app().log().debug("MailSettings event received: UseSslForSmtpFinished.");
        emit toggleUseSSLFinished();
        break;
    case MailSettingsEvent::kChangePortsFinished:
        app().log().debug("MailSettings event received: ChangePortsFinished.");
        emit changePortFinished();
        break;
    default:
        app().log().error("Unknown MailSettings event received.");
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processKeychainEvent(KeychainEvent const &event)
{
    switch (event.event_case())
    {
    case KeychainEvent::kChangeKeychainFinished:
        app().log().debug("Keychain event received: ChangeKeychainFinished.");
        emit changeKeychainFinished();
        break;
    case KeychainEvent::kHasNoKeychain:
        app().log().debug("Keychain event received: HasNoKeychain.");
        emit hasNoKeychain();
        break;
    case KeychainEvent::kRebuildKeychain:
        app().log().debug("Keychain event received: RebuildKeychain.");
        emit rebuildKeychain();
        break;
    default:
        app().log().error("Unknown Keychain event received.");
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processMailEvent(MailEvent const &event)
{
    switch (event.event_case())
    {
    case MailEvent::kNoActiveKeyForRecipientEvent:
        app().log().debug("Mail event received: kNoActiveKeyForRecipientEvent.");
        emit noActiveKeyForRecipient(QString::fromStdString(event.noactivekeyforrecipientevent().email()));
        break;
    case MailEvent::kAddressChanged:
        app().log().debug("Mail event received: AddressChanged.");
        emit addressChanged(QString::fromStdString(event.addresschanged().address()));
        break;
    case MailEvent::kAddressChangedLogout:
        app().log().debug("Mail event received: AddressChangedLogout.");
        emit addressChangedLogout(QString::fromStdString(event.addresschangedlogout().address()));
        break;
    case MailEvent::kApiCertIssue:
        emit apiCertIssue();
        app().log().debug("Mail event received: ApiCertIssue.");
        break;
    default:
        app().log().error("Unknown Mail event received.");
    }
}


//****************************************************************************************************************************************************
/// \param[in] event The event.
//****************************************************************************************************************************************************
void GRPCClient::processUserEvent(UserEvent const &event)
{
    switch (event.event_case())
    {
    case UserEvent::kToggleSplitModeFinished:
    {
        QString const userID = QString::fromStdString(event.togglesplitmodefinished().userid());
        app().log().debug(QString("User event received: ToggleSplitModeFinished (userID = %1).").arg(userID));
        emit toggleSplitModeFinished(userID);
        break;
    }
    case UserEvent::kUserDisconnected:
    {
        QString const username = QString::fromStdString(event.userdisconnected().username());
        app().log().debug(QString("User event received: UserDisconnected (username =  %1).").arg(username));
        emit userDisconnected(username);
        break;
    }
    case UserEvent::kUserChanged:
    {
        QString const userID = QString::fromStdString(event.userchanged().userid());
        app().log().debug(QString("User event received: UserChanged (userID = %1).").arg(userID));
        emit userChanged(userID);
        break;
    }
    default:
        app().log().error("Unknown User event received.");
    }
}
