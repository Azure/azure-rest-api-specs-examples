using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConnectedVMwarevSphere;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2020-10-01-preview/examples/CreateInventoryItem.json
// this example is just showing the usage of "InventoryItems_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InventoryItemResource created on azure
// for more information of creating InventoryItemResource, please refer to the document of InventoryItemResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vcenterName = "ContosoVCenter";
string inventoryItemName = "testItem";
ResourceIdentifier inventoryItemResourceId = InventoryItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vcenterName, inventoryItemName);
InventoryItemResource inventoryItem = client.GetInventoryItemResource(inventoryItemResourceId);

// invoke the operation
InventoryItemData data = new InventoryItemData(InventoryType.ResourcePool);
ArmOperation<InventoryItemResource> lro = await inventoryItem.UpdateAsync(WaitUntil.Completed, data);
InventoryItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
InventoryItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
