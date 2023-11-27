using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/VirtualNetworkScalableGatewayUpdate.json
// this example is just showing the usage of "VirtualNetworkGateways_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this VirtualNetworkGatewayResource
VirtualNetworkGatewayCollection collection = resourceGroupResource.GetVirtualNetworkGateways();

// invoke the operation
string virtualNetworkGatewayName = "ergw";
VirtualNetworkGatewayData data = new VirtualNetworkGatewayData()
{
    IPConfigurations =
    {
    new VirtualNetworkGatewayIPConfiguration()
    {
    PrivateIPAllocationMethod = NetworkIPAllocationMethod.Static,
    SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
    PublicIPAddressId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
    Name = "gwipconfig1",
    }
    },
    GatewayType = VirtualNetworkGatewayType.ExpressRoute,
    VpnType = VpnType.PolicyBased,
    EnableBgp = false,
    Active = false,
    DisableIPSecReplayProtection = false,
    Sku = new VirtualNetworkGatewaySku()
    {
        Name = VirtualNetworkGatewaySkuName.ErGwScale,
        Tier = VirtualNetworkGatewaySkuTier.ErGwScale,
    },
    VpnClientConfiguration = null,
    BgpSettings = null,
    NatRules =
    {
    new VirtualNetworkGatewayNatRuleData()
    {
    VpnNatRuleType = VpnNatRuleType.Static,
    Mode = VpnNatRuleMode.EgressSnat,
    InternalMappings =
    {
    new VpnNatRuleMapping()
    {
    AddressSpace = "10.10.0.0/24",
    }
    },
    ExternalMappings =
    {
    new VpnNatRuleMapping()
    {
    AddressSpace = "50.0.0.0/24",
    }
    },
    IPConfigurationId = "",
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule1"),
    Name = "natRule1",
    },new VirtualNetworkGatewayNatRuleData()
    {
    VpnNatRuleType = VpnNatRuleType.Static,
    Mode = VpnNatRuleMode.IngressSnat,
    InternalMappings =
    {
    new VpnNatRuleMapping()
    {
    AddressSpace = "20.10.0.0/24",
    }
    },
    ExternalMappings =
    {
    new VpnNatRuleMapping()
    {
    AddressSpace = "30.0.0.0/24",
    }
    },
    IPConfigurationId = "",
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule2"),
    Name = "natRule2",
    }
    },
    EnableBgpRouteTranslationForNat = false,
    AllowVirtualWanTraffic = false,
    AllowRemoteVnetTraffic = false,
    Location = new AzureLocation("centralus"),
};
ArmOperation<VirtualNetworkGatewayResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, virtualNetworkGatewayName, data);
VirtualNetworkGatewayResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualNetworkGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
