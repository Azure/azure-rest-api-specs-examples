using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-11-01/Schedulers_CreateOrUpdate.json
// this example is just showing the usage of "Scheduler_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "EE9BD735-67CE-4A90-89C4-439D3F6A4C93";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DurableTaskSchedulerResource
DurableTaskSchedulerCollection collection = resourceGroupResource.GetDurableTaskSchedulers();

// invoke the operation
string schedulerName = "testscheduler";
DurableTaskSchedulerData data = new DurableTaskSchedulerData(new AzureLocation("northcentralus"))
{
    Properties = new DurableTaskSchedulerProperties(new string[] { "10.0.0.0/8" }, new DurableTaskSchedulerSku(DurableTaskSchedulerSkuName.Dedicated)),
    Tags =
    {
    ["department"] = "research",
    ["development"] = "true"
    },
};
ArmOperation<DurableTaskSchedulerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, schedulerName, data);
DurableTaskSchedulerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DurableTaskSchedulerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
