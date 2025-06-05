using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/DeleteRaiBlocklistItems.json
// this example is just showing the usage of "RaiBlocklistItems_BatchDelete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RaiBlocklistResource created on azure
// for more information of creating RaiBlocklistResource, please refer to the document of RaiBlocklistResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string raiBlocklistName = "raiBlocklistName";
ResourceIdentifier raiBlocklistResourceId = RaiBlocklistResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, raiBlocklistName);
RaiBlocklistResource raiBlocklist = client.GetRaiBlocklistResource(raiBlocklistResourceId);

// invoke the operation
BinaryData raiBlocklistItemsNames = BinaryData.FromObjectAsJson(new object[]
{
"myblocklistitem1",
"myblocklistitem2"
});
await raiBlocklist.BatchDeleteRaiBlocklistItemAsync(raiBlocklistItemsNames);

Console.WriteLine("Succeeded");
