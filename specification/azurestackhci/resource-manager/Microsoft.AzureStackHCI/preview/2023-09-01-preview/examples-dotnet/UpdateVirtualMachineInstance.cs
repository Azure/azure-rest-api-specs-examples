using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Hci;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/UpdateVirtualMachineInstance.json
// this example is just showing the usage of "VirtualMachineInstances_Update" operation, for the dependent resources, they will have to be created separately.

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
VirtualMachineInstancePatch patch = new VirtualMachineInstancePatch()
{
    Properties = new VirtualMachineInstanceUpdateProperties()
    {
        StorageDataDisks =
        {
        new WritableSubResource()
        {
        Id = new ResourceIdentifier("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
        }
        },
    },
};
ArmOperation<VirtualMachineInstanceResource> lro = await virtualMachineInstance.UpdateAsync(WaitUntil.Completed, patch);
VirtualMachineInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
