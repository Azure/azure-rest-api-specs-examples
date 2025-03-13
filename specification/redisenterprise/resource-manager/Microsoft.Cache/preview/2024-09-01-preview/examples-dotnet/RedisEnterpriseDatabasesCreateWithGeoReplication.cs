using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise.Models;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseDatabasesCreateWithGeoReplication.json
// this example is just showing the usage of "Databases_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterpriseClusterResource created on azure
// for more information of creating RedisEnterpriseClusterResource, please refer to the document of RedisEnterpriseClusterResource
string subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
string resourceGroupName = "rg1";
string clusterName = "cache1";
ResourceIdentifier redisEnterpriseClusterResourceId = RedisEnterpriseClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
RedisEnterpriseClusterResource redisEnterpriseCluster = client.GetRedisEnterpriseClusterResource(redisEnterpriseClusterResourceId);

// get the collection of this RedisEnterpriseDatabaseResource
RedisEnterpriseDatabaseCollection collection = redisEnterpriseCluster.GetRedisEnterpriseDatabases();

// invoke the operation
string databaseName = "default";
RedisEnterpriseDatabaseData data = new RedisEnterpriseDatabaseData
{
    ClientProtocol = RedisEnterpriseClientProtocol.Encrypted,
    Port = 10000,
    ClusteringPolicy = RedisEnterpriseClusteringPolicy.EnterpriseCluster,
    EvictionPolicy = RedisEnterpriseEvictionPolicy.NoEviction,
    GeoReplication = new RedisEnterpriseDatabaseGeoReplication
    {
        GroupNickname = "groupName",
        LinkedDatabases = {new RedisEnterpriseLinkedDatabase
        {
        Id = new ResourceIdentifier("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
        }, new RedisEnterpriseLinkedDatabase
        {
        Id = new ResourceIdentifier("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8e/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
        }},
    },
    AccessKeysAuthentication = AccessKeysAuthentication.Enabled,
};
ArmOperation<RedisEnterpriseDatabaseResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, databaseName, data);
RedisEnterpriseDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisEnterpriseDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
