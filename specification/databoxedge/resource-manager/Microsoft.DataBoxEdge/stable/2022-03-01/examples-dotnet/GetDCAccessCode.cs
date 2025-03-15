using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataBoxEdge.Models;
using Azure.ResourceManager.DataBoxEdge;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/GetDCAccessCode.json
// this example is just showing the usage of "Orders_ListDCAccessCode" operation, for the dependent resources, they will have to be created separately.

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
DataBoxEdgeDataCenterAccessCode result = await dataBoxEdgeOrder.GetDataCenterAccessCodeAsync();

Console.WriteLine($"Succeeded: {result}");
