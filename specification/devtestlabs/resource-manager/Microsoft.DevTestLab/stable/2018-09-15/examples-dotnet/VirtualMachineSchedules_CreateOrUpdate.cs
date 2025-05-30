using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachineSchedules_CreateOrUpdate.json
// this example is just showing the usage of "VirtualMachineSchedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DevTestLabVmResource created on azure
// for more information of creating DevTestLabVmResource, please refer to the document of DevTestLabVmResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
string labName = "{labName}";
string vmName = "{vmName}";
ResourceIdentifier devTestLabVmResourceId = DevTestLabVmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, vmName);
DevTestLabVmResource devTestLabVm = client.GetDevTestLabVmResource(devTestLabVmResourceId);

// get the collection of this DevTestLabVmScheduleResource
DevTestLabVmScheduleCollection collection = devTestLabVm.GetDevTestLabVmSchedules();

// invoke the operation
string name = "LabVmsShutdown";
DevTestLabScheduleData data = new DevTestLabScheduleData(new AzureLocation("{location}"))
{
    Status = DevTestLabEnableStatus.Enabled,
    TaskType = "LabVmsShutdownTask",
    WeeklyRecurrence = new DevTestLabWeekDetails
    {
        Weekdays = { "Friday", "Saturday", "Sunday" },
        Time = "1700",
    },
    DailyRecurrenceTime = "1900",
    HourlyRecurrenceMinute = 30,
    TimeZoneId = "Pacific Standard Time",
    NotificationSettings = new DevTestLabNotificationSettings
    {
        Status = DevTestLabEnableStatus.Enabled,
        TimeInMinutes = 30,
        WebhookUri = new Uri("{webhookUrl}"),
        EmailRecipient = "{email}",
        NotificationLocale = "EN",
    },
    TargetResourceId = "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}",
    Tags =
    {
    ["tagName1"] = "tagValue1"
    },
};
ArmOperation<DevTestLabVmScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DevTestLabVmScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
