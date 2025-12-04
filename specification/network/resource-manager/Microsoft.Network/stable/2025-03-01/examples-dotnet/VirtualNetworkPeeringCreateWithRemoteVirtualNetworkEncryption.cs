using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkPeeringCreateWithRemoteVirtualNetworkEncryption.json
// this example is just showing the usage of "VirtualNetworkPeerings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkPeeringResource created on azure
// for more information of creating VirtualNetworkPeeringResource, please refer to the document of VirtualNetworkPeeringResource
string subscriptionId = "subid";
string resourceGroupName = "peerTest";
string virtualNetworkName = "vnet1";
string virtualNetworkPeeringName = "peer";
ResourceIdentifier virtualNetworkPeeringResourceId = VirtualNetworkPeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName, virtualNetworkPeeringName);
VirtualNetworkPeeringResource virtualNetworkPeering = client.GetVirtualNetworkPeeringResource(virtualNetworkPeeringResourceId);

// invoke the operation
VirtualNetworkPeeringData data = new VirtualNetworkPeeringData
{
    AllowVirtualNetworkAccess = true,
    AllowForwardedTraffic = true,
    AllowGatewayTransit = false,
    UseRemoteGateways = false,
    RemoteVirtualNetworkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/peerTest/providers/Microsoft.Network/virtualNetworks/vnet2"),
};
ArmOperation<VirtualNetworkPeeringResource> lro = await virtualNetworkPeering.UpdateAsync(WaitUntil.Completed, data);
VirtualNetworkPeeringResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualNetworkPeeringData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
