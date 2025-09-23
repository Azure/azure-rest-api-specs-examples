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

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/P2SVpnGatewayPut.json
// this example is just showing the usage of "P2sVpnGateways_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this P2SVpnGatewayResource
P2SVpnGatewayCollection collection = resourceGroupResource.GetP2SVpnGateways();

// invoke the operation
string gatewayName = "p2sVpnGateway1";
P2SVpnGatewayData data = new P2SVpnGatewayData
{
    VirtualHubId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
    P2SConnectionConfigurations = {new P2SConnectionConfiguration
    {
    VpnClientAddressPool = new VirtualNetworkAddressSpace
    {
    AddressPrefixes = {"101.3.0.0/16"},
    },
    RoutingConfiguration = new RoutingConfiguration
    {
    AssociatedRouteTableId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
    PropagatedRouteTables = new PropagatedRouteTable
    {
    Labels = {"label1", "label2"},
    Ids = {new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
    }, new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
    }, new WritableSubResource
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
    }},
    },
    VnetRoutes = new VnetRoute
    {
    StaticRoutes = {},
    },
    },
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
    Name = "P2SConnectionConfig1",
    }},
    VpnGatewayScaleUnit = 1,
    VpnServerConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"),
    CustomDnsServers = { "1.1.1.1", "2.2.2.2" },
    IsRoutingPreferenceInternet = false,
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<P2SVpnGatewayResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, gatewayName, data);
P2SVpnGatewayResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
P2SVpnGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
