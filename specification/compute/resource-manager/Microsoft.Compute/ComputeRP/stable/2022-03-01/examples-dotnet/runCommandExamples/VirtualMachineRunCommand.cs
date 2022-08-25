using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/runCommandExamples/VirtualMachineRunCommand.json
// this example is just showing the usage of "VirtualMachines_RunCommand" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this VirtualMachineResource created on azure
// for more information of creating VirtualMachineResource, please refer to the document of VirtualMachineResource
string subscriptionId = "24fb23e3-6ba3-41f0-9b6e-e41131d5d61e";
string resourceGroupName = "crptestar98131";
string vmName = "vm3036";
ResourceIdentifier virtualMachineResourceId = VirtualMachineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName);
VirtualMachineResource virtualMachine = client.GetVirtualMachineResource(virtualMachineResourceId);

// invoke the operation
RunCommandInput input = new RunCommandInput("RunPowerShellScript");
ArmOperation<VirtualMachineRunCommandResult> lro = await virtualMachine.RunCommandAsync(WaitUntil.Completed, input);
VirtualMachineRunCommandResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
