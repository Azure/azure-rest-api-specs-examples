const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a database in a RedisEnterprise cluster.
 *
 * @summary Gets information about a database in a RedisEnterprise cluster.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseDatabasesGet.json
 */
async function redisEnterpriseDatabasesGet() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.get(resourceGroupName, clusterName, databaseName);
  console.log(result);
}
