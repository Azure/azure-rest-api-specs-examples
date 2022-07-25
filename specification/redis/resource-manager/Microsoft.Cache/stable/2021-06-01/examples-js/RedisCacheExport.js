const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Export data from the redis cache to blobs in a container.
 *
 * @summary Export data from the redis cache to blobs in a container.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheExport.json
 */
async function redisCacheExport() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const parameters = {
    format: "RDB",
    container: "https://contosostorage.blob.core.window.net/urltoBlobContainer?sasKeyParameters",
    prefix: "datadump1",
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginExportDataAndWait(resourceGroupName, name, parameters);
  console.log(result);
}

redisCacheExport().catch(console.error);
