using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachine_Update_DetachDataDiskUsingToBeDetachedProperty.json
// this example is just showing the usage of "VirtualMachines_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string vmName = "myVM";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);

// invoke the operation
VirtualMachinePatch patch = new VirtualMachinePatch()
{
    HardwareProfile = new VirtualMachineHardwareProfile()
    {
        VmSize = VirtualMachineSizeType.StandardD2V2,
    },
    StorageProfile = new VirtualMachineStorageProfile()
    {
        ImageReference = new ImageReference()
        {
            Publisher = "MicrosoftWindowsServer",
            Offer = "WindowsServer",
            Sku = "2016-Datacenter",
            Version = "latest",
        },
        OSDisk = new VirtualMachineOSDisk(DiskCreateOptionType.FromImage)
        {
            Name = "myVMosdisk",
            Caching = CachingType.ReadWrite,
            ManagedDisk = new VirtualMachineManagedDisk()
            {
                StorageAccountType = StorageAccountType.StandardLrs,
            },
        },
        DataDisks =
        {
        new VirtualMachineDataDisk(0,DiskCreateOptionType.Empty)
        {
        DiskSizeGB = 1023,
        ToBeDetached = true,
        },new VirtualMachineDataDisk(1,DiskCreateOptionType.Empty)
        {
        DiskSizeGB = 1023,
        ToBeDetached = false,
        }
        },
    },
    OSProfile = new VirtualMachineOSProfile()
    {
        ComputerName = "myVM",
        AdminUsername = "{your-username}",
        AdminPassword = "{your-password}",
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
ArmOperation<VirtualMachineResource> lro = await virtualMachine.UpdateAsync(WaitUntil.Completed, patch);
VirtualMachineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
