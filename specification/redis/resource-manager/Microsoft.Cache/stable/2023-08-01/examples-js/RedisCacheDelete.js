const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Redis cache.
 *
 * @summary Deletes a Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheDelete.json
 */
async function redisCacheDelete() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const name = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginDeleteAndWait(resourceGroupName, name);
  console.log(result);
}
