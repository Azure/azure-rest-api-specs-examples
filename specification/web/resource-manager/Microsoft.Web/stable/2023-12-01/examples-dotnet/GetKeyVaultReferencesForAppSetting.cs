using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/GetKeyVaultReferencesForAppSetting.json
// this example is just showing the usage of "WebApps_GetAppSettingKeyVaultReference" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResource created on azure
// for more information of creating WebSiteResource, please refer to the document of WebSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "testc6282";
ResourceIdentifier webSiteResourceId = WebSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResource webSite = client.GetWebSiteResource(webSiteResourceId);

// get the collection of this SiteConfigAppsettingResource
SiteConfigAppsettingCollection collection = webSite.GetSiteConfigAppsettings();

// invoke the operation
string appSettingKey = "setting";
NullableResponse<SiteConfigAppsettingResource> response = await collection.GetIfExistsAsync(appSettingKey);
SiteConfigAppsettingResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiKeyVaultReferenceData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
