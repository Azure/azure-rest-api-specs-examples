const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the patching schedule of a redis cache.
 *
 * @summary Deletes the patching schedule of a redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCachePatchSchedulesDelete.json
 */
async function redisCachePatchSchedulesDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const defaultParam = "default";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.patchSchedules.delete(resourceGroupName, name, defaultParam);
  console.log(result);
}

redisCachePatchSchedulesDelete().catch(console.error);
