using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DurableTask.Models;
using Azure.ResourceManager.DurableTask;

// Generated from example definition: 2026-02-01/PrivateEndpointConnections_Get_MaximumSet_Gen.json
// this example is just showing the usage of "Schedulers_GetPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DurableTaskSchedulerResource created on azure
// for more information of creating DurableTaskSchedulerResource, please refer to the document of DurableTaskSchedulerResource
string subscriptionId = "851A7597-D699-45CC-899B-7487A5B3B775";
string resourceGroupName = "rgdurabletask";
string schedulerName = "testscheduler";
ResourceIdentifier durableTaskSchedulerResourceId = DurableTaskSchedulerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, schedulerName);
DurableTaskSchedulerResource durableTaskScheduler = client.GetDurableTaskSchedulerResource(durableTaskSchedulerResourceId);

// get the collection of this DurableTaskPrivateEndpointConnectionResource
DurableTaskPrivateEndpointConnectionCollection collection = durableTaskScheduler.GetPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "spzckqrbhfnabu";
NullableResponse<DurableTaskPrivateEndpointConnectionResource> response = await collection.GetIfExistsAsync(privateEndpointConnectionName);
DurableTaskPrivateEndpointConnectionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DurableTaskPrivateEndpointConnectionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
