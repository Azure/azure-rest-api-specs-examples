using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;
using Azure.ResourceManager.AppService.Models;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/ListSlotBackups.json
// this example is just showing the usage of "WebApps_ListSiteBackupsSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteSlotResource created on azure
// for more information of creating WebSiteSlotResource, please refer to the document of WebSiteSlotResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string name = "tests346";
string slot = "staging";
ResourceIdentifier webSiteSlotResourceId = WebSiteSlotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, slot);
WebSiteSlotResource webSiteSlot = client.GetWebSiteSlotResource(webSiteSlotResourceId);

// invoke the operation and iterate over the result
await foreach (WebAppBackupData item in webSiteSlot.GetAllSiteBackupSlotDataAsync())
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine($"Succeeded");
