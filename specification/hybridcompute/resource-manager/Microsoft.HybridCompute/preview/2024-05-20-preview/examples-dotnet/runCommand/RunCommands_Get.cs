using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridCompute.Models;
using Azure.ResourceManager.HybridCompute;

// Generated from example definition: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2024-05-20-preview/examples/runCommand/RunCommands_Get.json
// this example is just showing the usage of "MachineRunCommands_Get" operation, for the dependent resources, they will have to be created separately.

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
MachineRunCommandResource result = await machineRunCommand.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineRunCommandData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
