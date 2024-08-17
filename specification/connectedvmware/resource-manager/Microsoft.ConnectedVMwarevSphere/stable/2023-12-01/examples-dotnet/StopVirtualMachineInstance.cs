using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ConnectedVMwarevSphere;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-12-01/examples/StopVirtualMachineInstance.json
// this example is just showing the usage of "VirtualMachineInstances_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VMwareVmInstanceResource created on azure
// for more information of creating VMwareVmInstanceResource, please refer to the document of VMwareVmInstanceResource
string resourceUri = "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
ResourceIdentifier vMwareVmInstanceResourceId = VMwareVmInstanceResource.CreateResourceIdentifier(resourceUri);
VMwareVmInstanceResource vMwareVmInstance = client.GetVMwareVmInstanceResource(vMwareVmInstanceResourceId);

// invoke the operation
StopVirtualMachineContent content = new StopVirtualMachineContent()
{
    SkipShutdown = true,
};
await vMwareVmInstance.StopAsync(WaitUntil.Completed, content: content);

Console.WriteLine($"Succeeded");
