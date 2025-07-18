using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VpnGatewayPut.json
// this example is just showing the usage of "VpnGateways_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VpnGatewayResource
VpnGatewayCollection collection = resourceGroupResource.GetVpnGateways();

// invoke the operation
string gatewayName = "gateway1";
VpnGatewayData data = new VpnGatewayData
{
    VirtualHubId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
    Connections = {new VpnConnectionData
    {
    RemoteVpnSiteId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
    VpnLinkConnections = {new VpnSiteLinkConnectionData
    {
    VpnSiteLinkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
    VpnConnectionProtocolType = VirtualNetworkGatewayConnectionProtocol.IkeV2,
    ConnectionBandwidth = 200,
    SharedKey = "key",
    EgressNatRules = {new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/nat03"),
    }},
    Name = "Connection-Link1",
    }},
    Name = "vpnConnection1",
    }},
    BgpSettings = new BgpSettings
    {
        Asn = 65515L,
        PeerWeight = 0,
        BgpPeeringAddresses = {new NetworkIPConfigurationBgpPeeringAddress
        {
        IPConfigurationId = "Instance0",
        CustomBgpIPAddresses = {"169.254.21.5"},
        }, new NetworkIPConfigurationBgpPeeringAddress
        {
        IPConfigurationId = "Instance1",
        CustomBgpIPAddresses = {"169.254.21.10"},
        }},
    },
    EnableBgpRouteTranslationForNat = false,
    IsRoutingPreferenceInternet = false,
    NatRules = {new VpnGatewayNatRuleData
    {
    VpnNatRuleType = VpnNatRuleType.Static,
    Mode = VpnNatRuleMode.EgressSnat,
    InternalMappings = {new VpnNatRuleMapping
    {
    AddressSpace = "0.0.0.0/26",
    }},
    ExternalMappings = {new VpnNatRuleMapping
    {
    AddressSpace = "192.168.0.0/26",
    }},
    IPConfigurationId = "",
    Name = "nat03",
    }},
    Location = new AzureLocation("westcentralus"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<VpnGatewayResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, gatewayName, data);
VpnGatewayResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VpnGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
