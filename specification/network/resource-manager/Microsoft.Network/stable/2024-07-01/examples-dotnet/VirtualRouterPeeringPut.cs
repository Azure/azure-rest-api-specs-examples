using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualRouterPeeringPut.json
// this example is just showing the usage of "VirtualRouterPeerings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualRouterResource created on azure
// for more information of creating VirtualRouterResource, please refer to the document of VirtualRouterResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualRouterName = "virtualRouter";
ResourceIdentifier virtualRouterResourceId = VirtualRouterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualRouterName);
VirtualRouterResource virtualRouter = client.GetVirtualRouterResource(virtualRouterResourceId);

// get the collection of this VirtualRouterPeeringResource
VirtualRouterPeeringCollection collection = virtualRouter.GetVirtualRouterPeerings();

// invoke the operation
string peeringName = "peering1";
VirtualRouterPeeringData data = new VirtualRouterPeeringData
{
    PeerAsn = 20000L,
    PeerIP = "192.168.1.5",
};
ArmOperation<VirtualRouterPeeringResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, peeringName, data);
VirtualRouterPeeringResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualRouterPeeringData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
