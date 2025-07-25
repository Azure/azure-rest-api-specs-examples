using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RedisEnterprise.Models;
using Azure.ResourceManager.RedisEnterprise;

// Generated from example definition: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2025-04-01/examples/RedisEnterpriseDatabasesForceUnlink.json
// this example is just showing the usage of "Databases_ForceUnlink" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisEnterpriseDatabaseResource created on azure
// for more information of creating RedisEnterpriseDatabaseResource, please refer to the document of RedisEnterpriseDatabaseResource
string subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
string resourceGroupName = "rg1";
string clusterName = "cache1";
string databaseName = "default";
ResourceIdentifier redisEnterpriseDatabaseResourceId = RedisEnterpriseDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, databaseName);
RedisEnterpriseDatabaseResource redisEnterpriseDatabase = client.GetRedisEnterpriseDatabaseResource(redisEnterpriseDatabaseResourceId);

// invoke the operation
ForceUnlinkRedisEnterpriseDatabaseContent content = new ForceUnlinkRedisEnterpriseDatabaseContent(new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default") });
await redisEnterpriseDatabase.ForceUnlinkAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
