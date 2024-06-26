using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ArcScVmm;
using Azure.ResourceManager.ArcScVmm.Models;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteInventoryItem.json
// this example is just showing the usage of "InventoryItems_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InventoryItemResource created on azure
// for more information of creating InventoryItemResource, please refer to the document of InventoryItemResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vmmServerName = "ContosoVMMServer";
string inventoryItemName = "12345678-1234-1234-1234-123456789abc";
ResourceIdentifier inventoryItemResourceId = InventoryItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmmServerName, inventoryItemName);
InventoryItemResource inventoryItem = client.GetInventoryItemResource(inventoryItemResourceId);

// invoke the operation
await inventoryItem.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
