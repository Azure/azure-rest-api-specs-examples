const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the access keys for the RedisEnterprise database.
 *
 * @summary Retrieves the access keys for the RedisEnterprise database.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseDatabasesListKeys.json
 */
async function redisEnterpriseDatabasesListKeys() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.listKeys(resourceGroupName, clusterName, databaseName);
  console.log(result);
}
