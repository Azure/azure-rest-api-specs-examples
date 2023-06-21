const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all patch schedules in the specified redis cache (there is only one).
 *
 * @summary Gets all patch schedules in the specified redis cache (there is only one).
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCachePatchSchedulesList.json
 */
async function redisCachePatchSchedulesList() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const cacheName = "cache1";
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.patchSchedules.listByRedisResource(resourceGroupName, cacheName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
