using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRouteCircuitPeeringGet.json
// this example is just showing the usage of "ExpressRouteCircuitPeerings_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitResource created on azure
// for more information of creating ExpressRouteCircuitResource, please refer to the document of ExpressRouteCircuitResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string circuitName = "circuitName";
ResourceIdentifier expressRouteCircuitResourceId = ExpressRouteCircuitResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName);
ExpressRouteCircuitResource expressRouteCircuit = client.GetExpressRouteCircuitResource(expressRouteCircuitResourceId);

// get the collection of this ExpressRouteCircuitPeeringResource
ExpressRouteCircuitPeeringCollection collection = expressRouteCircuit.GetExpressRouteCircuitPeerings();

// invoke the operation
string peeringName = "MicrosoftPeering";
bool result = await collection.ExistsAsync(peeringName);

Console.WriteLine($"Succeeded: {result}");
