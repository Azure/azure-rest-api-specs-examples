const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the access keys for the RedisEnterprise database.
 *
 * @summary Retrieves the access keys for the RedisEnterprise database.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesListKeys.json
 */
async function redisEnterpriseDatabasesListKeys() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.listKeys(resourceGroupName, clusterName, databaseName);
  console.log(result);
}

redisEnterpriseDatabasesListKeys().catch(console.error);
