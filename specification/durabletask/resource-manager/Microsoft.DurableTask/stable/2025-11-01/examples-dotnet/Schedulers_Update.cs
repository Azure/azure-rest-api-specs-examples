using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2025-11-01/Schedulers_Update.json
// this example is just showing the usage of "Scheduler_Update" operation, for the dependent resources, they will have to be created separately.

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
DurableTaskSchedulerPatch patch = new DurableTaskSchedulerPatch
{
    Properties = new DurableTaskSchedulerPatchProperties
    {
        IPAllowlist = { "10.0.0.0/8" },
        Sku = new DurableTaskSchedulerSkuUpdate
        {
            Name = DurableTaskSchedulerSkuName.Dedicated,
            Capacity = 3,
        },
    },
    Tags =
    {
    ["hello"] = "world"
    },
};
ArmOperation<DurableTaskSchedulerResource> lro = await durableTaskScheduler.UpdateAsync(WaitUntil.Completed, patch);
DurableTaskSchedulerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DurableTaskSchedulerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
