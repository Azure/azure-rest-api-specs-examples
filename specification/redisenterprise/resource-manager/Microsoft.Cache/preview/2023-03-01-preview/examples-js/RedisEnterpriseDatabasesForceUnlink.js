const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Forcibly removes the link to the specified database resource.
 *
 * @summary Forcibly removes the link to the specified database resource.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseDatabasesForceUnlink.json
 */
async function howToUnlinkADatabaseDuringARegionalOutage() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const parameters = {
    ids: [
      "/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.beginForceUnlinkAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    parameters
  );
  console.log(result);
}
