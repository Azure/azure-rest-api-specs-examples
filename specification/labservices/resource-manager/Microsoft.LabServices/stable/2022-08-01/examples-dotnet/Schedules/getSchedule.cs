using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LabServices;
using Azure.ResourceManager.LabServices.Models;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Schedules/getSchedule.json
// this example is just showing the usage of "Schedules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LabServicesScheduleResource created on azure
// for more information of creating LabServicesScheduleResource, please refer to the document of LabServicesScheduleResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string labName = "testlab";
string scheduleName = "schedule1";
ResourceIdentifier labServicesScheduleResourceId = LabServicesScheduleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, scheduleName);
LabServicesScheduleResource labServicesSchedule = client.GetLabServicesScheduleResource(labServicesScheduleResourceId);

// invoke the operation
LabServicesScheduleResource result = await labServicesSchedule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LabServicesScheduleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
