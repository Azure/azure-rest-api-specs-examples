const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update an existing Redis cache.
 *
 * @summary Update an existing Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-05-01/examples/RedisCacheUpdate.json
 */
async function redisCacheUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const parameters = {
    enableNonSslPort: true,
    replicasPerPrimary: 2,
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginUpdateAndWait(resourceGroupName, name, parameters);
  console.log(result);
}

redisCacheUpdate().catch(console.error);
