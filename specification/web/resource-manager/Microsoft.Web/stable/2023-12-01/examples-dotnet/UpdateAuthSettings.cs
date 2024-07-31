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

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/UpdateAuthSettings.json
// this example is just showing the usage of "WebApps_UpdateAuthSettings" operation, for the dependent resources, they will have to be created separately.

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
SiteAuthSettings siteAuthSettings = new SiteAuthSettings()
{
    IsEnabled = true,
    RuntimeVersion = "~1",
    UnauthenticatedClientAction = UnauthenticatedClientAction.RedirectToLoginPage,
    IsTokenStoreEnabled = true,
    AllowedExternalRedirectUrls =
    {
    "sitef6141.customdomain.net","sitef6141.customdomain.info"
    },
    DefaultProvider = BuiltInAuthenticationProvider.Google,
    TokenRefreshExtensionHours = 120,
    ClientId = "42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com",
};
SiteAuthSettings result = await webSite.UpdateAuthSettingsAsync(siteAuthSettings);

Console.WriteLine($"Succeeded: {result}");
