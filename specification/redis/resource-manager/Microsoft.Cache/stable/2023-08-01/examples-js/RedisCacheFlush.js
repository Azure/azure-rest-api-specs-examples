const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes all of the keys in a cache.
 *
 * @summary Deletes all of the keys in a cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheFlush.json
 */
async function redisCacheFlush() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subcription-id";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "resource-group-name";
  const cacheName = "cache-name";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginFlushCacheAndWait(resourceGroupName, cacheName);
  console.log(result);
}
