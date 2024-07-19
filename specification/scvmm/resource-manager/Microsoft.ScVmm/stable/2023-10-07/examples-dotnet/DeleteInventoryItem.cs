using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteInventoryItem.json
// this example is just showing the usage of "InventoryItems_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmInventoryItemResource created on azure
// for more information of creating ScVmmInventoryItemResource, please refer to the document of ScVmmInventoryItemResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vmmServerName = "ContosoVMMServer";
string inventoryItemResourceName = "12345678-1234-1234-1234-123456789abc";
ResourceIdentifier scVmmInventoryItemResourceId = ScVmmInventoryItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmmServerName, inventoryItemResourceName);
ScVmmInventoryItemResource scVmmInventoryItem = client.GetScVmmInventoryItemResource(scVmmInventoryItemResourceId);

// invoke the operation
await scVmmInventoryItem.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
