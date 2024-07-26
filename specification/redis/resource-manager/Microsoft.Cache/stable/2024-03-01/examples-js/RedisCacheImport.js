const { RedisManagementClient } = require("@azure/arm-rediscache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Import data into Redis cache.
 *
 * @summary Import data into Redis cache.
 * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheImport.json
 */
async function redisCacheImport() {
  const subscriptionId = process.env["REDIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDIS_RESOURCE_GROUP"] || "rg1";
  const name = "cache1";
  const parameters = {
    format: "RDB",
    files: ["http://fileuris.contoso.com/pathtofile1"],
    storageSubscriptionId: "storageSubId",
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisManagementClient(credential, subscriptionId);
  const result = await client.redis.beginImportDataAndWait(resourceGroupName, name, parameters);
  console.log(result);
}
