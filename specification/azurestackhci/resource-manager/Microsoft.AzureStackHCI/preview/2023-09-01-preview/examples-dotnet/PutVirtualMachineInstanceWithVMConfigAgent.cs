using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualMachineInstanceWithVMConfigAgent.json
// this example is just showing the usage of "VirtualMachineInstances_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineInstanceResource created on azure
// for more information of creating VirtualMachineInstanceResource, please refer to the document of VirtualMachineInstanceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier virtualMachineInstanceResourceId = VirtualMachineInstanceResource.CreateResourceIdentifier(resourceUri);
VirtualMachineInstanceResource virtualMachineInstance = client.GetVirtualMachineInstanceResource(virtualMachineInstanceResourceId);

// invoke the operation
VirtualMachineInstanceData data = new VirtualMachineInstanceData()
{
    ExtendedLocation = new ArcVmExtendedLocation()
    {
        Name = "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location",
        ExtendedLocationType = ArcVmExtendedLocationType.CustomLocation,
    },
    HardwareProfile = new VirtualMachineInstancePropertiesHardwareProfile()
    {
        VmSize = VmSizeEnum.Default,
    },
    NetworkInterfaces =
    {
    new WritableSubResource()
    {
    Id = new ResourceIdentifier("test-nic"),
    }
    },
    OSProfile = new VirtualMachineInstancePropertiesOSProfile()
    {
        AdminPassword = "password",
        AdminUsername = "localadmin",
        ComputerName = "luamaster",
        WindowsConfiguration = new VirtualMachineInstancePropertiesOSProfileWindowsConfiguration()
        {
            ProvisionVmConfigAgent = true,
        },
    },
    SecurityProfile = new VirtualMachineInstancePropertiesSecurityProfile()
    {
        EnableTPM = true,
        SecureBootEnabled = true,
    },
    StorageProfile = new VirtualMachineInstancePropertiesStorageProfile()
    {
        ImageReferenceId = new ResourceIdentifier("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
        VmConfigStoragePathId = new ResourceIdentifier("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-container"),
    },
};
ArmOperation<VirtualMachineInstanceResource> lro = await virtualMachineInstance.CreateOrUpdateAsync(WaitUntil.Completed, data);
VirtualMachineInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
