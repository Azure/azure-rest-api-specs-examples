using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCircuitUpdateTags.json
// this example is just showing the usage of "ExpressRouteCircuits_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitResource created on azure
// for more information of creating ExpressRouteCircuitResource, please refer to the document of ExpressRouteCircuitResource
string subscriptionId = "subid";
string resourceGroupName = "ertest";
string circuitName = "er1";
ResourceIdentifier expressRouteCircuitResourceId = ExpressRouteCircuitResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName);
ExpressRouteCircuitResource expressRouteCircuit = client.GetExpressRouteCircuitResource(expressRouteCircuitResourceId);

// invoke the operation
NetworkTagsObject networkTagsObject = new NetworkTagsObject
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ExpressRouteCircuitResource result = await expressRouteCircuit.UpdateAsync(networkTagsObject);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteCircuitData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
