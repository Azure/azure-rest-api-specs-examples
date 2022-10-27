const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace the patching schedule for Redis cache.
 *
 * @summary Create or replace the patching schedule for Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCachePatchSchedulesCreateOrUpdate.json
 */
async function redisCachePatchSchedulesCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const defaultParam = "default";
  const parameters = {
    scheduleEntries: [
      { dayOfWeek: "Monday", maintenanceWindow: "PT5H", startHourUtc: 12 },
      { dayOfWeek: "Tuesday", startHourUtc: 12 },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.patchSchedules.createOrUpdate(
    resourceGroupName,
    name,
    defaultParam,
    parameters
  );
  console.log(result);
}

redisCachePatchSchedulesCreateOrUpdate().catch(console.error);
