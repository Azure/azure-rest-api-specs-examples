using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Vm.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Hci.Vm;

// Generated from example definition: 2025-06-01-preview/VirtualMachineInstances_CreateOrUpdate_Put_Virtual_Machine_Instance_With_Os_Disk.json
// this example is just showing the usage of "VirtualMachineInstance_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciVmInstanceResource created on azure
// for more information of creating HciVmInstanceResource, please refer to the document of HciVmInstanceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier hciVmInstanceResourceId = HciVmInstanceResource.CreateResourceIdentifier(resourceUri);
HciVmInstanceResource hciVmInstance = client.GetHciVmInstanceResource(hciVmInstanceResourceId);

// invoke the operation
HciVmInstanceData data = new HciVmInstanceData
{
    Properties = new HciVmInstanceProperties
    {
        HardwareProfile = new HciVmInstanceHardwareProfile
        {
            VmSize = HciVmSize.Default,
        },
        NetworkInterfaces = {new WritableSubResource
        {
        Id = new ResourceIdentifier("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/networkInterfaces/test-nic"),
        }},
        SecurityProfile = new HciVmInstanceSecurityProfile
        {
            IsTpmEnabled = true,
            SecureBootEnabled = true,
        },
        StorageProfile = new HciVmInstanceStorageProfile
        {
            OSDisk = new HciVmInstanceStorageProfileOSDisk
            {
                Id = new ResourceIdentifier("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
            },
            VmConfigStoragePathId = new ResourceIdentifier("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container"),
        },
    },
    ExtendedLocation = new HciVmExtendedLocation
    {
        Name = "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
        Type = HciVmExtendedLocationType.CustomLocation,
    },
};
ArmOperation<HciVmInstanceResource> lro = await hciVmInstance.CreateOrUpdateAsync(WaitUntil.Completed, data);
HciVmInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciVmInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
