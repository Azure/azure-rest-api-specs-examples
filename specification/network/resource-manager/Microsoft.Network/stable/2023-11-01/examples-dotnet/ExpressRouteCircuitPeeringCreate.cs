using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCircuitPeeringCreate.json
// this example is just showing the usage of "ExpressRouteCircuitPeerings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCircuitPeeringResource created on azure
// for more information of creating ExpressRouteCircuitPeeringResource, please refer to the document of ExpressRouteCircuitPeeringResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string circuitName = "circuitName";
string peeringName = "AzurePrivatePeering";
ResourceIdentifier expressRouteCircuitPeeringResourceId = ExpressRouteCircuitPeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, circuitName, peeringName);
ExpressRouteCircuitPeeringResource expressRouteCircuitPeering = client.GetExpressRouteCircuitPeeringResource(expressRouteCircuitPeeringResourceId);

// invoke the operation
ExpressRouteCircuitPeeringData data = new ExpressRouteCircuitPeeringData()
{
    PeerASN = 200,
    PrimaryPeerAddressPrefix = "192.168.16.252/30",
    SecondaryPeerAddressPrefix = "192.168.18.252/30",
    VlanId = 200,
};
ArmOperation<ExpressRouteCircuitPeeringResource> lro = await expressRouteCircuitPeering.UpdateAsync(WaitUntil.Completed, data);
ExpressRouteCircuitPeeringResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteCircuitPeeringData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
