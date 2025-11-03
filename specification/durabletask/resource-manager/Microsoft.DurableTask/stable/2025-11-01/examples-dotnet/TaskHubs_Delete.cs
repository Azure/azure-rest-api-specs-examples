using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-11-01/TaskHubs_Delete.json
// this example is just showing the usage of "TaskHub_Delete" operation, for the dependent resources, they will have to be created separately.

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
await durableTaskHub.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
