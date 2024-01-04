using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/runCommand/RunCommands_CreateOrUpdate.json
// this example is just showing the usage of "MachineRunCommands_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridComputeMachineResource created on azure
// for more information of creating HybridComputeMachineResource, please refer to the document of HybridComputeMachineResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
ResourceIdentifier hybridComputeMachineResourceId = HybridComputeMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName);
HybridComputeMachineResource hybridComputeMachine = client.GetHybridComputeMachineResource(hybridComputeMachineResourceId);

// get the collection of this MachineRunCommandResource
MachineRunCommandCollection collection = hybridComputeMachine.GetMachineRunCommands();

// invoke the operation
string runCommandName = "myRunCommand";
MachineRunCommandData data = new MachineRunCommandData(new AzureLocation("eastus2"))
{
    Source = new MachineRunCommandScriptSource()
    {
        Script = "Write-Host Hello World!",
    },
    Parameters =
    {
    new RunCommandInputParameter("param1","value1"),new RunCommandInputParameter("param2","value2")
    },
    AsyncExecution = false,
    RunAsUser = "user1",
    RunAsPassword = "<runAsPassword>",
    TimeoutInSeconds = 3600,
    OutputBlobUri = new Uri("https://mystorageaccount.blob.core.windows.net/myscriptoutputcontainer/MyScriptoutput.txt"),
    ErrorBlobUri = new Uri("https://mystorageaccount.blob.core.windows.net/mycontainer/MyScriptError.txt"),
};
ArmOperation<MachineRunCommandResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, runCommandName, data);
MachineRunCommandResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineRunCommandData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
