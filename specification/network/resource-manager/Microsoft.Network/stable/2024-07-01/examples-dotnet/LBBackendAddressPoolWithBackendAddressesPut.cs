using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/LBBackendAddressPoolWithBackendAddressesPut.json
// this example is just showing the usage of "LoadBalancerBackendAddressPools_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackendAddressPoolResource created on azure
// for more information of creating BackendAddressPoolResource, please refer to the document of BackendAddressPoolResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb";
string backendAddressPoolName = "backend";
ResourceIdentifier backendAddressPoolResourceId = BackendAddressPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName, backendAddressPoolName);
BackendAddressPoolResource backendAddressPool = client.GetBackendAddressPoolResource(backendAddressPoolResourceId);

// invoke the operation
BackendAddressPoolData data = new BackendAddressPoolData
{
    LoadBalancerBackendAddresses = {new LoadBalancerBackendAddress
    {
    Name = "address1",
    VirtualNetworkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
    IPAddress = "10.0.0.4",
    }, new LoadBalancerBackendAddress
    {
    Name = "address2",
    VirtualNetworkId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb"),
    IPAddress = "10.0.0.5",
    }},
};
ArmOperation<BackendAddressPoolResource> lro = await backendAddressPool.UpdateAsync(WaitUntil.Completed, data);
BackendAddressPoolResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackendAddressPoolData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
