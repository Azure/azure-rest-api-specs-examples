const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Flushes all the keys in this database and also from its linked databases.
 *
 * @summary Flushes all the keys in this database and also from its linked databases.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2025-05-01-preview/examples/RedisEnterpriseDatabasesFlush.json
 */
async function howToFlushAllTheKeysInTheDatabase() {
  const subscriptionId =
    process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const parameters = {
    ids: [
      "/subscriptions/e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default",
    ],
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.beginFlushAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    options,
  );
  console.log(result);
}
