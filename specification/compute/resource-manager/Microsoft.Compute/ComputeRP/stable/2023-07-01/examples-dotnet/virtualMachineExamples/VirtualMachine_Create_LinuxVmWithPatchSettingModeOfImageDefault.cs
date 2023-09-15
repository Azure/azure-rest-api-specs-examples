using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachine_Create_LinuxVmWithPatchSettingModeOfImageDefault.json
// this example is just showing the usage of "VirtualMachines_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this VirtualMachineResource
VirtualMachineCollection collection = resourceGroupResource.GetVirtualMachines();

// invoke the operation
string vmName = "myVM";
VirtualMachineData data = new VirtualMachineData(new AzureLocation("westus"))
{
    HardwareProfile = new VirtualMachineHardwareProfile()
    {
        VmSize = VirtualMachineSizeType.StandardD2SV3,
    },
    StorageProfile = new VirtualMachineStorageProfile()
    {
        ImageReference = new ImageReference()
        {
            Publisher = "Canonical",
            Offer = "UbuntuServer",
            Sku = "16.04-LTS",
            Version = "latest",
        },
        OSDisk = new VirtualMachineOSDisk(DiskCreateOptionType.FromImage)
        {
            Name = "myVMosdisk",
            Caching = CachingType.ReadWrite,
            ManagedDisk = new VirtualMachineManagedDisk()
            {
                StorageAccountType = StorageAccountType.PremiumLrs,
            },
        },
    },
    OSProfile = new VirtualMachineOSProfile()
    {
        ComputerName = "myVM",
        AdminUsername = "{your-username}",
        AdminPassword = "{your-password}",
        LinuxConfiguration = new LinuxConfiguration()
        {
            ProvisionVmAgent = true,
            PatchSettings = new LinuxPatchSettings()
            {
                PatchMode = LinuxVmGuestPatchMode.ImageDefault,
            },
        },
    },
    NetworkProfile = new VirtualMachineNetworkProfile()
    {
        NetworkInterfaces =
        {
        new VirtualMachineNetworkInterfaceReference()
        {
        Primary = true,
        Id = new ResourceIdentifier("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
        }
        },
    },
};
ArmOperation<VirtualMachineResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vmName, data);
VirtualMachineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
