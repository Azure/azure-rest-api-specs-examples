const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reboot specified Redis node(s). This operation requires write permission to the cache resource. There can be potential data loss.
 *
 * @summary Reboot specified Redis node(s). This operation requires write permission to the cache resource. There can be potential data loss.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheForceReboot.json
 */
async function redisCacheForceReboot() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const parameters = {
    ports: [13000, 15001],
    rebootType: "AllNodes",
    shardId: 0,
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.forceReboot(resourceGroupName, name, parameters);
  console.log(result);
}

redisCacheForceReboot().catch(console.error);
