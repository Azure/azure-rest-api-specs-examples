using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithServiceArtifactReference.json
// this example is just showing the usage of "VirtualMachineScaleSets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VirtualMachineScaleSetResource
VirtualMachineScaleSetCollection collection = resourceGroupResource.GetVirtualMachineScaleSets();

// invoke the operation
string virtualMachineScaleSetName = "{vmss-name}";
VirtualMachineScaleSetData data = new VirtualMachineScaleSetData(new AzureLocation("eastus2euap"))
{
    Sku = new ComputeSku()
    {
        Name = "Standard_A1",
        Tier = "Standard",
        Capacity = 3,
    },
    UpgradePolicy = new VirtualMachineScaleSetUpgradePolicy()
    {
        Mode = VirtualMachineScaleSetUpgradeMode.Automatic,
        AutomaticOSUpgradePolicy = new AutomaticOSUpgradePolicy()
        {
            EnableAutomaticOSUpgrade = true,
        },
    },
    VirtualMachineProfile = new VirtualMachineScaleSetVmProfile()
    {
        OSProfile = new VirtualMachineScaleSetOSProfile()
        {
            ComputerNamePrefix = "{vmss-name}",
            AdminUsername = "{your-username}",
            AdminPassword = "{your-password}",
        },
        StorageProfile = new VirtualMachineScaleSetStorageProfile()
        {
            ImageReference = new ImageReference()
            {
                Publisher = "MicrosoftWindowsServer",
                Offer = "WindowsServer",
                Sku = "2022-Datacenter",
                Version = "latest",
            },
            OSDisk = new VirtualMachineScaleSetOSDisk(DiskCreateOptionType.FromImage)
            {
                Name = "osDisk",
                Caching = CachingType.ReadWrite,
            },
        },
        NetworkProfile = new VirtualMachineScaleSetNetworkProfile()
        {
            NetworkInterfaceConfigurations =
            {
            new VirtualMachineScaleSetNetworkConfiguration("{vmss-name}")
            {
            Primary = true,
            IPConfigurations =
            {
            new VirtualMachineScaleSetIPConfiguration("{vmss-name}")
            {
            SubnetId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
            }
            },
            EnableIPForwarding = true,
            }
            },
        },
        ServiceArtifactReferenceId = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/serviceArtifacts/serviceArtifactName/vmArtifactsProfiles/vmArtifactsProfilesName"),
    },
    Overprovision = true,
};
ArmOperation<VirtualMachineScaleSetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, virtualMachineScaleSetName, data);
VirtualMachineScaleSetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineScaleSetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
