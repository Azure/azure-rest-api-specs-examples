const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a database
 *
 * @summary Creates a database
 * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/RedisEnterpriseDatabasesCreate.json
 */
async function redisEnterpriseDatabasesCreate() {
  const subscriptionId =
    process.env["REDISENTERPRISE_SUBSCRIPTION_ID"] || "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
  const resourceGroupName = process.env["REDISENTERPRISE_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cache1";
  const databaseName = "default";
  const parameters = {
    accessKeysAuthentication: "Enabled",
    clientProtocol: "Encrypted",
    clusteringPolicy: "EnterpriseCluster",
    deferUpgrade: "NotDeferred",
    evictionPolicy: "AllKeysLRU",
    modules: [
      { name: "RedisBloom", args: "ERROR_RATE 0.00 INITIAL_SIZE 400" },
      { name: "RedisTimeSeries", args: "RETENTION_POLICY 20" },
      { name: "RediSearch" },
    ],
    persistence: { aofEnabled: true, aofFrequency: "1s" },
    port: 10000,
  };
  const credential = new DefaultAzureCredential();
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    databaseName,
    parameters,
  );
  console.log(result);
}
