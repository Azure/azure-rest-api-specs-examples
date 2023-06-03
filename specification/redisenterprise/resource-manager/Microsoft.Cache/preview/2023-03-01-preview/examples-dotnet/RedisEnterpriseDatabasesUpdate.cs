using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RedisEnterprise;
using Azure.ResourceManager.RedisEnterprise.Models;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseDatabasesUpdate.json
// this example is just showing the usage of "Databases_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterpriseDatabaseResource created on azure
// for more information of creating RedisEnterpriseDatabaseResource, please refer to the document of RedisEnterpriseDatabaseResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string clusterName = "cache1";
string databaseName = "default";
ResourceIdentifier redisEnterpriseDatabaseResourceId = RedisEnterpriseDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
RedisEnterpriseDatabaseResource redisEnterpriseDatabase = client.GetRedisEnterpriseDatabaseResource(redisEnterpriseDatabaseResourceId);

// invoke the operation
RedisEnterpriseDatabasePatch patch = new RedisEnterpriseDatabasePatch()
{
    ClientProtocol = RedisEnterpriseClientProtocol.Encrypted,
    EvictionPolicy = RedisEnterpriseEvictionPolicy.AllKeysLru,
    Persistence = new RedisPersistenceSettings()
    {
        IsRdbEnabled = true,
        RdbFrequency = PersistenceSettingRdbFrequency.TwelveHours,
    },
};
ArmOperation<RedisEnterpriseDatabaseResource> lro = await redisEnterpriseDatabase.UpdateAsync(WaitUntil.Completed, patch);
RedisEnterpriseDatabaseResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RedisEnterpriseDatabaseData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
