using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PrivateLinkServiceCreate.json
// this example is just showing the usage of "PrivateLinkServices_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateLinkServiceResource created on azure
// for more information of creating PrivateLinkServiceResource, please refer to the document of PrivateLinkServiceResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string serviceName = "testPls";
ResourceIdentifier privateLinkServiceResourceId = PrivateLinkServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
PrivateLinkServiceResource privateLinkService = client.GetPrivateLinkServiceResource(privateLinkServiceResourceId);

// invoke the operation
PrivateLinkServiceData data = new PrivateLinkServiceData
{
    LoadBalancerFrontendIPConfigurations = {new FrontendIPConfigurationData
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
    }},
    IPConfigurations = {new PrivateLinkServiceIPConfiguration
    {
    PrivateIPAddress = "10.0.1.4",
    PrivateIPAllocationMethod = NetworkIPAllocationMethod.Static,
    Subnet = new SubnetData
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
    },
    PrivateIPAddressVersion = NetworkIPVersion.IPv4,
    Name = "fe-lb",
    }},
    VisibilitySubscriptions = { "subscription1", "subscription2", "subscription3" },
    AutoApprovalSubscriptions = { "subscription1", "subscription2" },
    Fqdns = { "fqdn1", "fqdn2", "fqdn3" },
    Location = new AzureLocation("eastus"),
};
ArmOperation<PrivateLinkServiceResource> lro = await privateLinkService.UpdateAsync(WaitUntil.Completed, data);
PrivateLinkServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateLinkServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
