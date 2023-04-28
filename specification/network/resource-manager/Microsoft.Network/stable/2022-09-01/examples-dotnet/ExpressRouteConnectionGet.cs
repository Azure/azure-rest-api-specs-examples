using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/ExpressRouteConnectionGet.json
// this example is just showing the usage of "ExpressRouteConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteGatewayResource created on azure
// for more information of creating ExpressRouteGatewayResource, please refer to the document of ExpressRouteGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "resourceGroupName";
string expressRouteGatewayName = "expressRouteGatewayName";
ResourceIdentifier expressRouteGatewayResourceId = ExpressRouteGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, expressRouteGatewayName);
ExpressRouteGatewayResource expressRouteGateway = client.GetExpressRouteGatewayResource(expressRouteGatewayResourceId);

// get the collection of this ExpressRouteConnectionResource
ExpressRouteConnectionCollection collection = expressRouteGateway.GetExpressRouteConnections();

// invoke the operation
string connectionName = "connectionName";
bool result = await collection.ExistsAsync(connectionName);

Console.WriteLine($"Succeeded: {result}");
