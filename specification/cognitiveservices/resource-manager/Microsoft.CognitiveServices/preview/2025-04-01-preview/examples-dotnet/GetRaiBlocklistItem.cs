using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CognitiveServices.Models;
using Azure.ResourceManager.CognitiveServices;

// Generated from example definition: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/GetRaiBlocklistItem.json
// this example is just showing the usage of "RaiBlocklistItems_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RaiBlocklistItemResource created on azure
// for more information of creating RaiBlocklistItemResource, please refer to the document of RaiBlocklistItemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroupName";
string accountName = "accountName";
string raiBlocklistName = "raiBlocklistName";
string raiBlocklistItemName = "raiBlocklistItemName";
ResourceIdentifier raiBlocklistItemResourceId = RaiBlocklistItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, raiBlocklistName, raiBlocklistItemName);
RaiBlocklistItemResource raiBlocklistItem = client.GetRaiBlocklistItemResource(raiBlocklistItemResourceId);

// invoke the operation
RaiBlocklistItemResource result = await raiBlocklistItem.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RaiBlocklistItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
