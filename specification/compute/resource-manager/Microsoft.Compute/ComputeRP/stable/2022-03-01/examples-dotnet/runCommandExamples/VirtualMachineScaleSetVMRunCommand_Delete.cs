using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/runCommandExamples/VirtualMachineScaleSetVMRunCommand_Delete.json
// this example is just showing the usage of "VirtualMachineScaleSetVMRunCommands_Delete" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this VirtualMachineScaleSetVmRunCommandResource created on azure
// for more information of creating VirtualMachineScaleSetVmRunCommandResource, please refer to the document of VirtualMachineScaleSetVmRunCommandResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string virtualMachineScaleSetName = "myvmScaleSet";
string instanceId = "0";
string runCommandName = "myRunCommand";
ResourceIdentifier virtualMachineScaleSetVmRunCommandResourceId = VirtualMachineScaleSetVmRunCommandResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualMachineScaleSetName, instanceId, runCommandName);
VirtualMachineScaleSetVmRunCommandResource virtualMachineScaleSetVmRunCommand = client.GetVirtualMachineScaleSetVmRunCommandResource(virtualMachineScaleSetVmRunCommandResourceId);
            
// invoke the operation
await virtualMachineScaleSetVmRunCommand.DeleteAsync(WaitUntil.Completed);
            
Console.WriteLine($"Succeeded");
