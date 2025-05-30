using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-07-31-preview/examples/gateway/Gateway_Delete.json
// this example is just showing the usage of "Gateways_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArcGatewayResource created on azure
// for more information of creating ArcGatewayResource, please refer to the document of ArcGatewayResource
string subscriptionId = "ffd506c8-3415-42d3-9612-fdb423fb17df";
string resourceGroupName = "myResourceGroup";
string gatewayName = "{gatewayName}";
ResourceIdentifier arcGatewayResourceId = ArcGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName);
ArcGatewayResource arcGateway = client.GetArcGatewayResource(arcGatewayResourceId);

// invoke the operation
await arcGateway.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
