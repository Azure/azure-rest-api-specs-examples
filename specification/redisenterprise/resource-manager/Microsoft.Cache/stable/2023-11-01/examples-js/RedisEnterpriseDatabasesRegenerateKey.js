const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the RedisEnterprise database's access keys.
 *
 * @summary Regenerates the RedisEnterprise database's access keys.
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseDatabasesRegenerateKey.json
 */
async function redisEnterpriseDatabasesRegenerateKey() {
  const subscriptionId = process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const parameters = { keyType: "Primary" };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.beginRegenerateKeyAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    parameters,
  );
  console.log(result);
}
