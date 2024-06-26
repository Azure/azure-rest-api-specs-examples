using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/ExpressRouteCircuitConnectionGet.json
// this example is just showing the usage of "ExpressRouteCircuitConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitPeeringResource created on azure
// for more information of creating ExpressRouteCircuitPeeringResource, please refer to the document of ExpressRouteCircuitPeeringResource
string subscriptionId = "subid1";
string resourceGroupName = "rg1";
string circuitName = "ExpressRouteARMCircuitA";
string peeringName = "AzurePrivatePeering";
ResourceIdentifier expressRouteCircuitPeeringResourceId = ExpressRouteCircuitPeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName, peeringName);
ExpressRouteCircuitPeeringResource expressRouteCircuitPeering = client.GetExpressRouteCircuitPeeringResource(expressRouteCircuitPeeringResourceId);

// get the collection of this ExpressRouteCircuitConnectionResource
ExpressRouteCircuitConnectionCollection collection = expressRouteCircuitPeering.GetExpressRouteCircuitConnections();

// invoke the operation
string connectionName = "circuitConnectionUSAUS";
bool result = await collection.ExistsAsync(connectionName);

Console.WriteLine($"Succeeded: {result}");
