using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_Get.json
// this example is just showing the usage of "VirtualMachineSchedules_Get" operation, for the dependent resources, they will have to be created separately.

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
DevTestLabVmScheduleResource result = await devTestLabVmSchedule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
