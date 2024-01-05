using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridCompute;
using Azure.ResourceManager.HybridCompute.Models;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-10-03-preview/examples/runCommand/RunCommands_Delete.json
// this example is just showing the usage of "MachineRunCommands_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineRunCommandResource created on azure
// for more information of creating MachineRunCommandResource, please refer to the document of MachineRunCommandResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string machineName = "myMachine";
string runCommandName = "myRunCommand";
ResourceIdentifier machineRunCommandResourceId = MachineRunCommandResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, machineName, runCommandName);
MachineRunCommandResource machineRunCommand = client.GetMachineRunCommandResource(machineRunCommandResourceId);

// invoke the operation
await machineRunCommand.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
