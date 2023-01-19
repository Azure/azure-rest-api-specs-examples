using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EdgeOrder;
using Azure.ResourceManager.EdgeOrder.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/DeleteOrderItemByName.json
// this example is just showing the usage of "DeleteOrderItemByName" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeOrderItemResource created on azure
// for more information of creating EdgeOrderItemResource, please refer to the document of EdgeOrderItemResource
string subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
string resourceGroupName = "TestRG";
string orderItemName = "TestOrderItemName01";
ResourceIdentifier edgeOrderItemResourceId = EdgeOrderItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, orderItemName);
EdgeOrderItemResource edgeOrderItem = client.GetEdgeOrderItemResource(edgeOrderItemResourceId);

// invoke the operation
await edgeOrderItem.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
