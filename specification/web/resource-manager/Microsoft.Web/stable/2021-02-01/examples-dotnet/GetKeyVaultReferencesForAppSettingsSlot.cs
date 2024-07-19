using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/GetKeyVaultReferencesForAppSettingsSlot.json
// this example is just showing the usage of "WebApps_GetAppSettingsKeyVaultReferencesSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteSlotResource created on azure
// for more information of creating WebSiteSlotResource, please refer to the document of WebSiteSlotResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "testc6282";
string slot = "stage";
ResourceIdentifier webSiteSlotResourceId = WebSiteSlotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, slot);
WebSiteSlotResource webSiteSlot = client.GetWebSiteSlotResource(webSiteSlotResourceId);

// get the collection of this WebSiteSlotConfigAppSettingResource
WebSiteSlotConfigAppSettingCollection collection = webSiteSlot.GetWebSiteSlotConfigAppSettings();

// invoke the operation and iterate over the result
await foreach (WebSiteSlotConfigAppSettingResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApiKeyVaultReferenceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
