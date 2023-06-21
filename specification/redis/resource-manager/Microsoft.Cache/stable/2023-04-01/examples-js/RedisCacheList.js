const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Redis caches in the specified subscription.
 *
 * @summary Gets all Redis caches in the specified subscription.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheList.json
 */
async function redisCacheList() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.redis.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
