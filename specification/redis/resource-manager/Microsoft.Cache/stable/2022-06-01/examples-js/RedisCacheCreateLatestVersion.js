const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace (overwrite/recreate, with potential downtime) an existing Redis cache.
 *
 * @summary Create or replace (overwrite/recreate, with potential downtime) an existing Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheCreateLatestVersion.json
 */
async function redisCacheCreateLatestVersion() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const parameters = {
    enableNonSslPort: true,
    location: "West US",
    minimumTlsVersion: "1.2",
    redisConfiguration: { maxmemoryPolicy: "allkeys-lru" },
    redisVersion: "Latest",
    replicasPerPrimary: 2,
    shardCount: 2,
    sku: { name: "Premium", capacity: 1, family: "P" },
    staticIP: "192.168.0.5",
    subnetId:
      "/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/network1/subnets/subnet1",
    zones: ["1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginCreateAndWait(resourceGroupName, name, parameters);
  console.log(result);
}

redisCacheCreateLatestVersion().catch(console.error);
