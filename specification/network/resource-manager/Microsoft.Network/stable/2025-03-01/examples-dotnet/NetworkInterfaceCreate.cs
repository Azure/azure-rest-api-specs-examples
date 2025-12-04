using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkInterfaceCreate.json
// this example is just showing the usage of "NetworkInterfaces_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this NetworkInterfaceResource
NetworkInterfaceCollection collection = resourceGroupResource.GetNetworkInterfaces();

// invoke the operation
string networkInterfaceName = "test-nic";
NetworkInterfaceData data = new NetworkInterfaceData
{
    IPConfigurations = {new NetworkInterfaceIPConfigurationData
    {
    Subnet = new SubnetData
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/rg1-vnet/subnets/default"),
    },
    PublicIPAddress = new PublicIPAddressData
    {
    Id = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"),
    },
    Name = "ipconfig1",
    }, new NetworkInterfaceIPConfigurationData
    {
    PrivateIPAddressPrefixLength = 28,
    Name = "ipconfig2",
    }},
    EnableAcceleratedNetworking = true,
    DisableTcpStateTracking = true,
    Location = new AzureLocation("eastus"),
};
ArmOperation<NetworkInterfaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkInterfaceName, data);
NetworkInterfaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkInterfaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
