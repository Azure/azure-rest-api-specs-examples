using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCircuitConnectionDelete.json
// this example is just showing the usage of "ExpressRouteCircuitConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitConnectionResource created on azure
// for more information of creating ExpressRouteCircuitConnectionResource, please refer to the document of ExpressRouteCircuitConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string circuitName = "ExpressRouteARMCircuitA";
string peeringName = "AzurePrivatePeering";
string connectionName = "circuitConnectionUSAUS";
ResourceIdentifier expressRouteCircuitConnectionResourceId = ExpressRouteCircuitConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName, peeringName, connectionName);
ExpressRouteCircuitConnectionResource expressRouteCircuitConnection = client.GetExpressRouteCircuitConnectionResource(expressRouteCircuitConnectionResourceId);

// invoke the operation
await expressRouteCircuitConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
