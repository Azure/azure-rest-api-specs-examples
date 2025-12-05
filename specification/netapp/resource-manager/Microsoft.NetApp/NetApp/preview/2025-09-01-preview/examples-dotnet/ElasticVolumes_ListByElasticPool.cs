using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/ElasticVolumes_ListByElasticPool.json
// this example is just showing the usage of "ElasticVolumes_ListByElasticPool" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticCapacityPoolResource created on azure
// for more information of creating ElasticCapacityPoolResource, please refer to the document of ElasticCapacityPoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
ResourceIdentifier elasticCapacityPoolResourceId = ElasticCapacityPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName);
ElasticCapacityPoolResource elasticCapacityPool = client.GetElasticCapacityPoolResource(elasticCapacityPoolResourceId);

// get the collection of this ElasticVolumeResource
ElasticVolumeCollection collection = elasticCapacityPool.GetElasticVolumes();

// invoke the operation and iterate over the result
await foreach (ElasticVolumeResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ElasticVolumeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
