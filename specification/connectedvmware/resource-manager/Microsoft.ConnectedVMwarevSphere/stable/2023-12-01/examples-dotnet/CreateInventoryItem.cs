using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.ConnectedVMwarevSphere;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-12-01/examples/CreateInventoryItem.json
// this example is just showing the usage of "InventoryItems_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VCenterInventoryItemResource created on azure
// for more information of creating VCenterInventoryItemResource, please refer to the document of VCenterInventoryItemResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vcenterName = "ContosoVCenter";
string inventoryItemName = "testItem";
ResourceIdentifier vCenterInventoryItemResourceId = VCenterInventoryItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vcenterName, inventoryItemName);
VCenterInventoryItemResource vCenterInventoryItem = client.GetVCenterInventoryItemResource(vCenterInventoryItemResourceId);

// invoke the operation
VCenterInventoryItemData data = new VCenterInventoryItemData(VCenterInventoryType.ResourcePool);
ArmOperation<VCenterInventoryItemResource> lro = await vCenterInventoryItem.UpdateAsync(WaitUntil.Completed, data);
VCenterInventoryItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VCenterInventoryItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
