using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Execute.json
// this example is just showing the usage of "VirtualMachineSchedules_Execute" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabVmScheduleResource created on azure
// for more information of creating DevTestLabVmScheduleResource, please refer to the document of DevTestLabVmScheduleResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string vmName = "{vmName}";
string name = "LabVmsShutdown";
ResourceIdentifier devTestLabVmScheduleResourceId = DevTestLabVmScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, vmName, name);
DevTestLabVmScheduleResource devTestLabVmSchedule = client.GetDevTestLabVmScheduleResource(devTestLabVmScheduleResourceId);

// invoke the operation
await devTestLabVmSchedule.ExecuteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
