using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/runCommandExamples/VirtualMachineRunCommand_CreateOrUpdate.json
// this example is just showing the usage of "VirtualMachineRunCommands_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this VirtualMachineRunCommandResource
VirtualMachineRunCommandCollection collection = virtualMachine.GetVirtualMachineRunCommands();

// invoke the operation
string runCommandName = "myRunCommand";
VirtualMachineRunCommandData data = new VirtualMachineRunCommandData(new AzureLocation("West US"))
{
    Source = new VirtualMachineRunCommandScriptSource
    {
        ScriptUri = new Uri("https://mystorageaccount.blob.core.windows.net/scriptcontainer/scriptURI"),
    },
    Parameters = { new RunCommandInputParameter("param1", "value1"), new RunCommandInputParameter("param2", "value2") },
    AsyncExecution = false,
    RunAsUser = "user1",
    RunAsPassword = "<runAsPassword>",
    TimeoutInSeconds = 3600,
    OutputBlobUri = new Uri("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
    ErrorBlobUri = new Uri("https://mystorageaccount.blob.core.windows.net/scriptcontainer/scriptURI"),
    OutputBlobManagedIdentity = new RunCommandManagedIdentity
    {
        ClientId = "22d35efb-0c99-4041-8c5b-6d24db33a69a",
    },
    TreatFailureAsDeploymentFailure = false,
};
ArmOperation<VirtualMachineRunCommandResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, runCommandName, data);
VirtualMachineRunCommandResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineRunCommandData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
