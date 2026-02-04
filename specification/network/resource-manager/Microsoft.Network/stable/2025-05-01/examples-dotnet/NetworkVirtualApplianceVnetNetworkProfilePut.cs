using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkVirtualApplianceVnetNetworkProfilePut.json
// this example is just showing the usage of "NetworkVirtualAppliances_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this NetworkVirtualApplianceResource
NetworkVirtualApplianceCollection collection = resourceGroupResource.GetNetworkVirtualAppliances();

// invoke the operation
string networkVirtualApplianceName = "nva";
NetworkVirtualApplianceData data = new NetworkVirtualApplianceData
{
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1")] = new UserAssignedIdentity()
        },
    },
    NvaSku = new VirtualApplianceSkuProperties
    {
        Vendor = "Cisco SDWAN",
        BundledScaleUnit = "1",
        MarketPlaceVersion = "latest",
    },
    BootStrapConfigurationBlobs = { "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrbootstrapconfig" },
    CloudInitConfigurationBlobs = { "https://csrncvhdstorage1.blob.core.windows.net/csrncvhdstoragecont/csrcloudinitconfig" },
    VirtualApplianceAsn = 10000L,
    NetworkInterfaceConfigurations = {new VirtualApplianceNetworkInterfaceConfiguration
    {
    NicType = NicTypeInRequest.PublicNic,
    VirtualApplianceNetworkInterfaceIPConfigurations = {new VirtualApplianceIPConfiguration
    {
    Name = "myPrimaryPublicIpConfig",
    IsPrimary = true,
    }, new VirtualApplianceIPConfiguration
    {
    Name = "myNonPrimaryPublicIpConfig",
    IsPrimary = false,
    }},
    }, new VirtualApplianceNetworkInterfaceConfiguration
    {
    NicType = NicTypeInRequest.PrivateNic,
    VirtualApplianceNetworkInterfaceIPConfigurations = {new VirtualApplianceIPConfiguration
    {
    Name = "myPrimaryPrivateIpConfig",
    IsPrimary = true,
    }, new VirtualApplianceIPConfiguration
    {
    Name = "myNonPrimaryPrivateIpConfig",
    IsPrimary = false,
    }},
    }},
    NvaInterfaceConfigurations = {new NvaInterfaceConfigurationsProperties
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
    PropertiesType = {NvaNicType.PrivateNic},
    Name = "dataInterface",
    }, new NvaInterfaceConfigurationsProperties
    {
    SubnetId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet2"),
    PropertiesType = {NvaNicType.PublicNic},
    Name = "managementInterface",
    }},
    Location = new AzureLocation("West US"),
    Tags =
    {
    ["key1"] = "value1"
    },
};
ArmOperation<NetworkVirtualApplianceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, networkVirtualApplianceName, data);
NetworkVirtualApplianceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkVirtualApplianceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
