using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/examples/PeerExpressRouteCircuitConnectionGet.json
// this example is just showing the usage of "PeerExpressRouteCircuitConnections_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this PeerExpressRouteCircuitConnectionResource
PeerExpressRouteCircuitConnectionCollection collection = expressRouteCircuitPeering.GetPeerExpressRouteCircuitConnections();

// invoke the operation
string connectionName = "60aee347-e889-4a42-8c1b-0aae8b1e4013";
bool result = await collection.ExistsAsync(connectionName);

Console.WriteLine($"Succeeded: {result}");
