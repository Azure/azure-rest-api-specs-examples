using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-11-01/TaskHubs_Get.json
// this example is just showing the usage of "TaskHub_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DurableTaskSchedulerResource created on azure
// for more information of creating DurableTaskSchedulerResource, please refer to the document of DurableTaskSchedulerResource
string subscriptionId = "EE9BD735-67CE-4A90-89C4-439D3F6A4C93";
string resourceGroupName = "rgopenapi";
string schedulerName = "testscheduler";
ResourceIdentifier durableTaskSchedulerResourceId = DurableTaskSchedulerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schedulerName);
DurableTaskSchedulerResource durableTaskScheduler = client.GetDurableTaskSchedulerResource(durableTaskSchedulerResourceId);

// get the collection of this DurableTaskHubResource
DurableTaskHubCollection collection = durableTaskScheduler.GetDurableTaskHubs();

// invoke the operation
string taskHubName = "testtaskhub";
NullableResponse<DurableTaskHubResource> response = await collection.GetIfExistsAsync(taskHubName);
DurableTaskHubResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DurableTaskHubData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
