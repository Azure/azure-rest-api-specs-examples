using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ScVmm.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/DeleteCheckpointVirtualMachineInstance.json
// this example is just showing the usage of "VirtualMachineInstances_DeleteCheckpoint" operation, for the dependent resources, they will have to be created separately.

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
VirtualMachineDeleteCheckpointContent content = new VirtualMachineDeleteCheckpointContent()
{
    Id = "Demo CheckpointID",
};
await scVmmVirtualMachineInstance.DeleteCheckpointAsync(WaitUntil.Completed, content: content);

Console.WriteLine($"Succeeded");
