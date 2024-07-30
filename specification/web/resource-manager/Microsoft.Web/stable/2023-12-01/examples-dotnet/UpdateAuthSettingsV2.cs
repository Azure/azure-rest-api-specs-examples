using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/UpdateAuthSettingsV2.json
// this example is just showing the usage of "WebApps_UpdateAuthSettingsV2" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResource created on azure
// for more information of creating WebSiteResource, please refer to the document of WebSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "sitef6141";
ResourceIdentifier webSiteResourceId = WebSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResource webSite = client.GetWebSiteResource(webSiteResourceId);

// invoke the operation
SiteAuthSettingsV2 siteAuthSettingsV2 = new SiteAuthSettingsV2()
{
    Platform = new AuthPlatform()
    {
        IsEnabled = true,
        RuntimeVersion = "~1",
        ConfigFilePath = "/auth/config.json",
    },
    GlobalValidation = new GlobalValidation()
    {
        IsAuthenticationRequired = true,
        UnauthenticatedClientAction = UnauthenticatedClientActionV2.Return403,
        ExcludedPaths =
        {
        "/nosecrets/Path"
        },
    },
    IdentityProviders = new AppServiceIdentityProviders()
    {
        Google = new AppServiceGoogleProvider()
        {
            IsEnabled = true,
            Registration = new ClientRegistration()
            {
                ClientId = "42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com",
                ClientSecretSettingName = "ClientSecret",
            },
            LoginScopes =
            {
            "admin"
            },
            ValidationAllowedAudiences =
            {
            "https://example.com"
            },
        },
    },
    Login = new WebAppLoginInfo()
    {
        RoutesLogoutEndpoint = "https://app.com/logout",
        TokenStore = new AppServiceTokenStore()
        {
            IsEnabled = true,
            TokenRefreshExtensionHours = 96,
            FileSystemDirectory = "/wwwroot/sites/example",
        },
        PreserveUrlFragmentsForLogins = true,
        AllowedExternalRedirectUrls =
        {
        "https://someurl.com"
        },
        CookieExpiration = new WebAppCookieExpiration()
        {
            Convention = CookieExpirationConvention.IdentityProviderDerived,
            TimeToExpiration = "2022:09-01T00:00Z",
        },
        Nonce = new LoginFlowNonceSettings()
        {
            ValidateNonce = true,
        },
    },
    HttpSettings = new AppServiceHttpSettings()
    {
        IsHttpsRequired = true,
        RoutesApiPrefix = "/authv2/",
        ForwardProxy = new AppServiceForwardProxy()
        {
            Convention = ForwardProxyConvention.Standard,
            CustomHostHeaderName = "authHeader",
            CustomProtoHeaderName = "customProtoHeader",
        },
    },
};
SiteAuthSettingsV2 result = await webSite.UpdateAuthSettingsV2Async(siteAuthSettingsV2);

Console.WriteLine($"Succeeded: {result}");
