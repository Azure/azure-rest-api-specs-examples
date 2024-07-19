using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCircuitRouteTableList.json
// this example is just showing the usage of "ExpressRouteCircuits_ListRoutesTable" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitPeeringResource created on azure
// for more information of creating ExpressRouteCircuitPeeringResource, please refer to the document of ExpressRouteCircuitPeeringResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string circuitName = "circuitName";
string peeringName = "peeringName";
ResourceIdentifier expressRouteCircuitPeeringResourceId = ExpressRouteCircuitPeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName, peeringName);
ExpressRouteCircuitPeeringResource expressRouteCircuitPeering = client.GetExpressRouteCircuitPeeringResource(expressRouteCircuitPeeringResourceId);

// invoke the operation
string devicePath = "devicePath";
ArmOperation<ExpressRouteCircuitsRoutesTableListResult> lro = await expressRouteCircuitPeering.GetRoutesTableExpressRouteCircuitAsync(WaitUntil.Completed, devicePath);
ExpressRouteCircuitsRoutesTableListResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
