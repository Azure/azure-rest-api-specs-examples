using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/UpdateVirtualMachineInstance.json
// this example is just showing the usage of "VirtualMachineInstances_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScVmmVirtualMachineInstanceResource created on azure
// for more information of creating ScVmmVirtualMachineInstanceResource, please refer to the document of ScVmmVirtualMachineInstanceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier scVmmVirtualMachineInstanceResourceId = ScVmmVirtualMachineInstanceResource.CreateResourceIdentifier(resourceUri);
ScVmmVirtualMachineInstanceResource scVmmVirtualMachineInstance = client.GetScVmmVirtualMachineInstanceResource(scVmmVirtualMachineInstanceResourceId);

// invoke the operation
ScVmmVirtualMachineInstancePatch patch = new ScVmmVirtualMachineInstancePatch
{
    HardwareProfile = new ScVmmHardwareProfileUpdate
    {
        MemoryMB = 4196,
        CpuCount = 4,
    },
};
ArmOperation<ScVmmVirtualMachineInstanceResource> lro = await scVmmVirtualMachineInstance.UpdateAsync(WaitUntil.Completed, patch);
ScVmmVirtualMachineInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScVmmVirtualMachineInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
