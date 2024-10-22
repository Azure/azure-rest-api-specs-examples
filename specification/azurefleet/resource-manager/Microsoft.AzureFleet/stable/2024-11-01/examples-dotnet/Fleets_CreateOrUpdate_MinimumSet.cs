using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeFleet.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ComputeFleet;

// Generated from example definition: 2024-11-01/Fleets_CreateOrUpdate_MinimumSet.json
// this example is just showing the usage of "Fleet_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1DC2F28C-A625-4B0E-9748-9885A3C9E9EB";
string resourceGroupName = "rgazurefleet";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ComputeFleetResource
ComputeFleetCollection collection = resourceGroupResource.GetComputeFleets();

// invoke the operation
string fleetName = "testFleet";
ComputeFleetData data = new ComputeFleetData(new AzureLocation("eastus2euap"))
{
    Properties = new ComputeFleetProperties(new ComputeFleetVmSizeProfile[]
{
new ComputeFleetVmSizeProfile("Standard_D2s_v3"),new ComputeFleetVmSizeProfile("Standard_D4s_v3"),new ComputeFleetVmSizeProfile("Standard_E2s_v3")
}, new ComputeFleetComputeProfile(new ComputeFleetVmProfile()
{
    OSProfile = new ComputeFleetVmssOSProfile()
    {
        ComputerNamePrefix = "prefix",
        AdminUsername = "azureuser",
        AdminPassword = "TestPassword$0",
        LinuxConfiguration = new ComputeFleetLinuxConfiguration()
        {
            IsPasswordAuthenticationDisabled = false,
        },
    },
    StorageProfile = new ComputeFleetVmssStorageProfile()
    {
        ImageReference = new ComputeFleetImageReference()
        {
            Publisher = "canonical",
            Offer = "0001-com-ubuntu-server-focal",
            Sku = "20_04-lts-gen2",
            Version = "latest",
        },
        OSDisk = new ComputeFleetVmssOSDisk(ComputeFleetDiskCreateOptionType.FromImage)
        {
            Caching = ComputeFleetCachingType.ReadWrite,
            OSType = ComputeFleetOperatingSystemType.Linux,
            ManagedDisk = new ComputeFleetVmssManagedDisk()
            {
                StorageAccountType = ComputeFleetStorageAccountType.StandardLrs,
            },
        },
    },
    NetworkProfile = new ComputeFleetVmssNetworkProfile()
    {
        NetworkInterfaceConfigurations =
{
new ComputeFleetVmssNetworkConfiguration("vmNameTest")
{
Properties = new ComputeFleetVmssNetworkConfigurationProperties(new ComputeFleetVmssIPConfiguration[]
{
new ComputeFleetVmssIPConfiguration("vmNameTest")
{
Properties = new ComputeFleetVmssIPConfigurationProperties()
{
SubnetId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"),
IsPrimary = true,
LoadBalancerBackendAddressPools =
{
new WritableSubResource()
{
Id = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"),
}
},
},
}
})
{
IsPrimary = true,
IsAcceleratedNetworkingEnabled = false,
IsIPForwardingEnabled = true,
},
}
},
    },
})
{
    ComputeApiVersion = "2023-09-01",
    PlatformFaultDomainCount = 1,
})
    {
        SpotPriorityProfile = new SpotPriorityProfile()
        {
            Capacity = 2,
            MinCapacity = 1,
            EvictionPolicy = ComputeFleetEvictionPolicy.Delete,
            AllocationStrategy = SpotAllocationStrategy.PriceCapacityOptimized,
            IsMaintainEnabled = true,
        },
        RegularPriorityProfile = new RegularPriorityProfile()
        {
            Capacity = 2,
            MinCapacity = 1,
            AllocationStrategy = RegularPriorityAllocationStrategy.LowestPrice,
        },
    },
    Tags =
    {
    ["key"] = "fleets-test",
    },
};
ArmOperation<ComputeFleetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, fleetName, data);
ComputeFleetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ComputeFleetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
