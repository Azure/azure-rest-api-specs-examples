using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.LabServices;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Schedules/putSchedule.json
// this example is just showing the usage of "Schedules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LabResource created on azure
// for more information of creating LabResource, please refer to the document of LabResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string labName = "testlab";
ResourceIdentifier labResourceId = LabResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName);
LabResource lab = client.GetLabResource(labResourceId);

// get the collection of this LabServicesScheduleResource
LabServicesScheduleCollection collection = lab.GetLabServicesSchedules();

// invoke the operation
string scheduleName = "schedule1";
LabServicesScheduleData data = new LabServicesScheduleData
{
    StartOn = DateTimeOffset.Parse("2020-05-26T12:00:00Z"),
    StopOn = DateTimeOffset.Parse("2020-05-26T18:00:00Z"),
    RecurrencePattern = new LabServicesRecurrencePattern(LabServicesRecurrenceFrequency.Daily, DateTimeOffset.Parse("2020-08-14T23:59:59Z"))
    {
        Interval = 2,
    },
    TimeZoneId = "America/Los_Angeles",
    Notes = BinaryData.FromObjectAsJson("Schedule 1 for students"),
};
ArmOperation<LabServicesScheduleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scheduleName, data);
LabServicesScheduleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LabServicesScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
