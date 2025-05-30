using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CostManagement.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.CostManagement;

// Generated from example definition: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2023-03-01/examples/scheduledActions/scheduledAction-createOrUpdate-private.json
// this example is just showing the usage of "ScheduledActions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this TenantScheduledActionResource
TenantScheduledActionCollection collection = tenantResource.GetTenantScheduledActions();

// invoke the operation
string name = "monthlyCostByResource";
ScheduledActionData data = new ScheduledActionData
{
    DisplayName = "Monthly Cost By Resource",
    Notification = new NotificationProperties(new string[] { "user@gmail.com", "team@gmail.com" }, "Cost by resource this month"),
    Schedule = new ScheduleProperties(ScheduleFrequency.Monthly, DateTimeOffset.Parse("2020-06-19T22:21:51.1287144Z"), DateTimeOffset.Parse("2021-06-19T22:21:51.1287144Z"))
    {
        HourOfDay = 10,
        DaysOfWeek = { ScheduledActionDaysOfWeek.Monday },
        WeeksOfMonth = { ScheduledActionWeeksOfMonth.First, ScheduledActionWeeksOfMonth.Third },
    },
    Status = ScheduledActionStatus.Enabled,
    ViewId = new ResourceIdentifier("/providers/Microsoft.CostManagement/views/swaggerExample"),
    Kind = ScheduledActionKind.Email,
};
string ifMatch = "";
ArmOperation<TenantScheduledActionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data, ifMatch: ifMatch);
TenantScheduledActionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScheduledActionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
