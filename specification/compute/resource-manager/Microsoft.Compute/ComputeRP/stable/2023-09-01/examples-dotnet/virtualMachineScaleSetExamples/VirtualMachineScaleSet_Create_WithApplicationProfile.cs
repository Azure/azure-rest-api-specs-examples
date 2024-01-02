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

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithApplicationProfile.json
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
VirtualMachineScaleSetData data = new VirtualMachineScaleSetData(new AzureLocation("westus"))
{
    Sku = new ComputeSku()
    {
        Name = "Standard_D1_v2",
        Tier = "Standard",
        Capacity = 3,
    },
    UpgradePolicy = new VirtualMachineScaleSetUpgradePolicy()
    {
        Mode = VirtualMachineScaleSetUpgradeMode.Manual,
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
                Sku = "2016-Datacenter",
                Version = "latest",
            },
            OSDisk = new VirtualMachineScaleSetOSDisk(DiskCreateOptionType.FromImage)
            {
                Caching = CachingType.ReadWrite,
                ManagedDisk = new VirtualMachineScaleSetManagedDisk()
                {
                    StorageAccountType = StorageAccountType.StandardLrs,
                },
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
        GalleryApplications =
        {
        new VirtualMachineGalleryApplication("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0")
        {
        Tags = "myTag1",
        Order = 1,
        ConfigurationReference = "https://mystorageaccount.blob.core.windows.net/configurations/settings.config",
        TreatFailureAsDeploymentFailure = true,
        EnableAutomaticUpgrade = false,
        },new VirtualMachineGalleryApplication("/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1")
        },
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
