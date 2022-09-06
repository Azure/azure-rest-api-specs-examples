const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerate Redis cache's access keys. This operation requires write permission to the cache resource.
 *
 * @summary Regenerate Redis cache's access keys. This operation requires write permission to the cache resource.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-05-01/examples/RedisCacheRegenerateKey.json
 */
async function redisCacheRegenerateKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const parameters = { keyType: "Primary" };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.regenerateKey(resourceGroupName, name, parameters);
  console.log(result);
}

redisCacheRegenerateKey().catch(console.error);
