using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/runCommandExamples/VirtualMachineRunCommand_Delete.json
// this example is just showing the usage of "VirtualMachineRunCommands_Delete" operation, for the dependent resources, they will have to be created separately.

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
await virtualMachineRunCommand.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
