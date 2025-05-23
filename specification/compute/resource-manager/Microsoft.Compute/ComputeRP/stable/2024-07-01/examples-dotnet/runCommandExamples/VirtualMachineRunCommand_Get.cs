using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/runCommandExamples/VirtualMachineRunCommand_Get.json
// this example is just showing the usage of "VirtualMachineRunCommands_GetByVirtualMachine" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualMachineRunCommandResource created on azure
// for more information of creating VirtualMachineRunCommandResource, please refer to the document of VirtualMachineRunCommandResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string vmName = "myVM";
string runCommandName = "myRunCommand";
ResourceIdentifier virtualMachineRunCommandResourceId = VirtualMachineRunCommandResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vmName, runCommandName);
VirtualMachineRunCommandResource virtualMachineRunCommand = client.GetVirtualMachineRunCommandResource(virtualMachineRunCommandResourceId);

// invoke the operation
VirtualMachineRunCommandResource result = await virtualMachineRunCommand.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualMachineRunCommandData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
