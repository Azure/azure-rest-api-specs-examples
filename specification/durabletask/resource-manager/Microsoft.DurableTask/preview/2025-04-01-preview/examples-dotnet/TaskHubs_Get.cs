using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-04-01-preview/TaskHubs_Get.json
// this example is just showing the usage of "TaskHub_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DurableTaskHubResource created on azure
// for more information of creating DurableTaskHubResource, please refer to the document of DurableTaskHubResource
string subscriptionId = "EE9BD735-67CE-4A90-89C4-439D3F6A4C93";
string resourceGroupName = "rgopenapi";
string schedulerName = "testscheduler";
string taskHubName = "testtaskhub";
ResourceIdentifier durableTaskHubResourceId = DurableTaskHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schedulerName, taskHubName);
DurableTaskHubResource durableTaskHub = client.GetDurableTaskHubResource(durableTaskHubResourceId);

// invoke the operation
DurableTaskHubResource result = await durableTaskHub.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DurableTaskHubData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
