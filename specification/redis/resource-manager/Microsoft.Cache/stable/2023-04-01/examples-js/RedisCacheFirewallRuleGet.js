const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a single firewall rule in a specified redis cache.
 *
 * @summary Gets a single firewall rule in a specified redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheFirewallRuleGet.json
 */
async function redisCacheFirewallRuleGet() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const cacheName = "cache1";
  const ruleName = "rule1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.get(resourceGroupName, cacheName, ruleName);
  console.log(result);
}
