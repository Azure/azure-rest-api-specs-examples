using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/runCommandExamples/VirtualMachineRunCommand_CreateOrUpdate.json
// this example is just showing the usage of "VirtualMachineRunCommands_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
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
    Source = new VirtualMachineRunCommandScriptSource()
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
};
ArmOperation<VirtualMachineRunCommandResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, runCommandName, data);
VirtualMachineRunCommandResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineRunCommandData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
