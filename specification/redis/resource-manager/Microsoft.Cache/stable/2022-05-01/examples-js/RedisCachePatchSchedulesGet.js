const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the patching schedule of a redis cache.
 *
 * @summary Gets the patching schedule of a redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-05-01/examples/RedisCachePatchSchedulesGet.json
 */
async function redisCachePatchSchedulesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.patchSchedules.get(resourceGroupName, name, defaultParam);
  console.log(result);
}

redisCachePatchSchedulesGet().catch(console.error);
