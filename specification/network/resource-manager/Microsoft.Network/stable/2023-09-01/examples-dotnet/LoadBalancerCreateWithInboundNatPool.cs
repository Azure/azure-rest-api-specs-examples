using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/LoadBalancerCreateWithInboundNatPool.json
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
    Zones =
    {
    },
    PrivateIPAllocationMethod = NetworkIPAllocationMethod.Dynamic,
    Subnet = new SubnetData()
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet"),
    },
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
    Name = "test",
    }
    },
    BackendAddressPools =
    {
    },
    LoadBalancingRules =
    {
    },
    Probes =
    {
    },
    InboundNatRules =
    {
    },
    InboundNatPools =
    {
    new LoadBalancerInboundNatPool()
    {
    FrontendIPConfigurationId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
    Protocol = LoadBalancingTransportProtocol.Tcp,
    FrontendPortRangeStart = 8080,
    FrontendPortRangeEnd = 8085,
    BackendPort = 8888,
    IdleTimeoutInMinutes = 10,
    EnableFloatingIP = true,
    EnableTcpReset = true,
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test"),
    Name = "test",
    }
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
