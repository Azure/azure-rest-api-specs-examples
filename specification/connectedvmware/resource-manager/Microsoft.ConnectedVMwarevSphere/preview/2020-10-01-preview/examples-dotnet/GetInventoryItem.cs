using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConnectedVMwarevSphere;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2020-10-01-preview/examples/GetInventoryItem.json
// this example is just showing the usage of "InventoryItems_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VCenterResource created on azure
// for more information of creating VCenterResource, please refer to the document of VCenterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vcenterName = "ContosoVCenter";
ResourceIdentifier vCenterResourceId = VCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vcenterName);
VCenterResource vCenter = client.GetVCenterResource(vCenterResourceId);

// get the collection of this InventoryItemResource
InventoryItemCollection collection = vCenter.GetInventoryItems();

// invoke the operation
string inventoryItemName = "testItem";
bool result = await collection.ExistsAsync(inventoryItemName);

Console.WriteLine($"Succeeded: {result}");
