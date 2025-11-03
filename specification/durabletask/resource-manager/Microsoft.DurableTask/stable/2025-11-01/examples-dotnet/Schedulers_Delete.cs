using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-11-01/Schedulers_Delete.json
// this example is just showing the usage of "Scheduler_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
await durableTaskScheduler.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
