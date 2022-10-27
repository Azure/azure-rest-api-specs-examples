const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a redis cache firewall rule
 *
 * @summary Create or update a redis cache firewall rule
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheFirewallRuleCreate.json
 */
async function redisCacheFirewallRuleCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const cacheName = "cache1";
  const ruleName = "rule1";
  const parameters = {
    endIP: "192.168.1.4",
    startIP: "192.168.1.1",
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.createOrUpdate(
    resourceGroupName,
    cacheName,
    ruleName,
    parameters
  );
  console.log(result);
}

redisCacheFirewallRuleCreate().catch(console.error);
