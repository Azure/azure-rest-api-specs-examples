using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBoxEdge;
using Azure.ResourceManager.DataBoxEdge.Models;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/OrderDelete.json
// this example is just showing the usage of "Orders_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeOrderResource created on azure
// for more information of creating DataBoxEdgeOrderResource, please refer to the document of DataBoxEdgeOrderResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
ResourceIdentifier dataBoxEdgeOrderResourceId = DataBoxEdgeOrderResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName);
DataBoxEdgeOrderResource dataBoxEdgeOrder = client.GetDataBoxEdgeOrderResource(dataBoxEdgeOrderResourceId);

// invoke the operation
await dataBoxEdgeOrder.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
