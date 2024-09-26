using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise.Models;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseDatabasesForceLink.json
// this example is just showing the usage of "Databases_ForceLinkToReplicationGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterpriseDatabaseResource created on azure
// for more information of creating RedisEnterpriseDatabaseResource, please refer to the document of RedisEnterpriseDatabaseResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string clusterName = "cache1";
string databaseName = "default";
ResourceIdentifier redisEnterpriseDatabaseResourceId = RedisEnterpriseDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
RedisEnterpriseDatabaseResource redisEnterpriseDatabase = client.GetRedisEnterpriseDatabaseResource(redisEnterpriseDatabaseResourceId);

// invoke the operation
ForceLinkContent content = new ForceLinkContent("groupName", new RedisEnterpriseLinkedDatabase[]
{
new RedisEnterpriseLinkedDatabase()
{
Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
},new RedisEnterpriseLinkedDatabase()
{
Id = new ResourceIdentifier("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
}
});
await redisEnterpriseDatabase.ForceLinkToReplicationGroupAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
