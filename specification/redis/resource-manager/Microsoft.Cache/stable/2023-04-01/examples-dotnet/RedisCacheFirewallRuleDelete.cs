using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Redis;

// Generated from example definition: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheFirewallRuleDelete.json
// this example is just showing the usage of "FirewallRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RedisFirewallRuleResource created on azure
// for more information of creating RedisFirewallRuleResource, please refer to the document of RedisFirewallRuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string cacheName = "cache1";
string ruleName = "rule1";
ResourceIdentifier redisFirewallRuleResourceId = RedisFirewallRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, cacheName, ruleName);
RedisFirewallRuleResource redisFirewallRule = client.GetRedisFirewallRuleResource(redisFirewallRuleResourceId);

// invoke the operation
await redisFirewallRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
