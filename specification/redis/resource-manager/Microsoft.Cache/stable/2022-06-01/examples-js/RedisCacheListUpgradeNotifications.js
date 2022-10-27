const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets any upgrade notifications for a Redis cache.
 *
 * @summary Gets any upgrade notifications for a Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheListUpgradeNotifications.json
 */
async function redisCacheListUpgradeNotifications() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const history = 5000;
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.redis.listUpgradeNotifications(resourceGroupName, name, history)) {
    resArray.push(item);
  }
  console.log(resArray);
}

redisCacheListUpgradeNotifications().catch(console.error);
