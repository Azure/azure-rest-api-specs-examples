const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Redis cache (resource description).
 *
 * @summary Gets a Redis cache (resource description).
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheGet.json
 */
async function redisCacheGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.get(resourceGroupName, name);
  console.log(result);
}

redisCacheGet().catch(console.error);
