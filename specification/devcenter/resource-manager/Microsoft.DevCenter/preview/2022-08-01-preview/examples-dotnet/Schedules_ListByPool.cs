using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DevCenter;
using Azure.ResourceManager.DevCenter.Models;

// Generated from example definition: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Schedules_ListByPool.json
// this example is just showing the usage of "Schedules_ListByPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PoolResource created on azure
// for more information of creating PoolResource, please refer to the document of PoolResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "rg1";
string projectName = "TestProject";
string poolName = "DevPool";
ResourceIdentifier poolResourceId = PoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, projectName, poolName);
PoolResource pool = client.GetPoolResource(poolResourceId);

// get the collection of this ScheduleResource
ScheduleCollection collection = pool.GetSchedules();

// invoke the operation and iterate over the result
await foreach (ScheduleResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ScheduleData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
