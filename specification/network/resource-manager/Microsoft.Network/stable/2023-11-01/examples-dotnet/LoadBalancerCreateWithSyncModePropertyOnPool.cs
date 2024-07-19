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

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/LoadBalancerCreateWithSyncModePropertyOnPool.json
// this example is just showing the usage of "LoadBalancers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this LoadBalancerResource
LoadBalancerCollection collection = resourceGroupResource.GetLoadBalancers();

// invoke the operation
string loadBalancerName = "lb";
LoadBalancerData data = new LoadBalancerData()
{
    Sku = new LoadBalancerSku()
    {
        Name = LoadBalancerSkuName.Standard,
    },
    FrontendIPConfigurations =
    {
    new FrontendIPConfigurationData()
    {
    Subnet = new SubnetData()
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
    },
    Name = "fe-lb",
    }
    },
    BackendAddressPools =
    {
    new BackendAddressPoolData()
    {
    VirtualNetworkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
    SyncMode = BackendAddressSyncMode.Automatic,
    Name = "be-lb",
    }
    },
    LoadBalancingRules =
    {
    new LoadBalancingRuleData()
    {
    FrontendIPConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
    BackendAddressPoolId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
    ProbeId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
    Protocol = LoadBalancingTransportProtocol.Tcp,
    LoadDistribution = LoadDistribution.Default,
    FrontendPort = 80,
    BackendPort = 80,
    IdleTimeoutInMinutes = 15,
    EnableFloatingIP = true,
    Name = "rulelb",
    }
    },
    Probes =
    {
    new ProbeData()
    {
    Protocol = ProbeProtocol.Http,
    Port = 80,
    IntervalInSeconds = 15,
    NumberOfProbes = 2,
    ProbeThreshold = 1,
    RequestPath = "healthcheck.aspx",
    Name = "probe-lb",
    }
    },
    InboundNatRules =
    {
    new InboundNatRuleData()
    {
    FrontendIPConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
    Protocol = LoadBalancingTransportProtocol.Tcp,
    FrontendPort = 3389,
    BackendPort = 3389,
    IdleTimeoutInMinutes = 15,
    EnableFloatingIP = true,
    Name = "in-nat-rule",
    }
    },
    InboundNatPools =
    {
    },
    OutboundRules =
    {
    },
    Location = new AzureLocation("eastus"),
};
ArmOperation<LoadBalancerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, loadBalancerName, data);
LoadBalancerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LoadBalancerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
