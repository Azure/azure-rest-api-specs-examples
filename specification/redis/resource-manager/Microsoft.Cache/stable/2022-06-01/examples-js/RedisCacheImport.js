const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Import data into Redis cache.
 *
 * @summary Import data into Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheImport.json
 */
async function redisCacheImport() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const name = "cache1";
  const parameters = {
    format: "RDB",
    files: ["http://fileuris.contoso.com/pathtofile1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginImportDataAndWait(resourceGroupName, name, parameters);
  console.log(result);
}

redisCacheImport().catch(console.error);
