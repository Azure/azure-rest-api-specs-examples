using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DevTestLabs.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DevTestLabs;

// Generated from example definition: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_CreateOrUpdate.json
// this example is just showing the usage of "GlobalSchedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DevTestLabGlobalScheduleResource
DevTestLabGlobalScheduleCollection collection = resourceGroupResource.GetDevTestLabGlobalSchedules();

// invoke the operation
string name = "labvmautostart";
DevTestLabScheduleData data = new DevTestLabScheduleData(default)
{
    Status = DevTestLabEnableStatus.Enabled,
    TaskType = "LabVmsStartupTask",
    WeeklyRecurrence = new DevTestLabWeekDetails
    {
        Weekdays = { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday" },
        Time = "0700",
    },
    TimeZoneId = "Hawaiian Standard Time",
};
ArmOperation<DevTestLabGlobalScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DevTestLabGlobalScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DevTestLabScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
