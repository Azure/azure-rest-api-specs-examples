using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2021-02-01/examples/VpnConnectionPut.json
// this example is just showing the usage of "VpnConnections_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnConnectionResource created on azure
// for more information of creating VpnConnectionResource, please refer to the document of VpnConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string gatewayName = "gateway1";
string connectionName = "vpnConnection1";
ResourceIdentifier vpnConnectionResourceId = VpnConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName, connectionName);
VpnConnectionResource vpnConnection = client.GetVpnConnectionResource(vpnConnectionResourceId);

// invoke the operation
VpnConnectionData data = new VpnConnectionData()
{
    RemoteVpnSiteId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
    TrafficSelectorPolicies =
    {
    },
    VpnLinkConnections =
    {
    new VpnSiteLinkConnectionData()
    {
    VpnSiteLinkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
    VpnLinkConnectionMode = VpnLinkConnectionMode.Default,
    VpnConnectionProtocolType = VirtualNetworkGatewayConnectionProtocol.IkeV2,
    ConnectionBandwidth = 200,
    SharedKey = "key",
    UsePolicyBasedTrafficSelectors = false,
    Name = "Connection-Link1",
    }
    },
};
ArmOperation<VpnConnectionResource> lro = await vpnConnection.UpdateAsync(WaitUntil.Completed, data);
VpnConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VpnConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
