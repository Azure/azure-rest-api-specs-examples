const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a database
 *
 * @summary creates a database
 * x-ms-original-file: 2026-02-01-preview/RedisEnterpriseDatabasesCreate.json
 */
async function redisEnterpriseDatabasesCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.create("rg1", "cache1", "default", {
    accessKeysAuthentication: "Enabled",
    clientProtocol: "Encrypted",
    clusteringPolicy: "EnterpriseCluster",
    deferUpgrade: "NotDeferred",
    evictionPolicy: "AllKeysLRU",
    notifyKeyspaceEvents: "KEA",
    modules: [
      { name: "RedisBloom", args: "ERROR_RATE 0.00 INITIAL_SIZE 400" },
      { name: "RedisTimeSeries", args: "RETENTION_POLICY 20" },
      { name: "RediSearch" },
    ],
    persistence: { aofEnabled: true, aofFrequency: "1s" },
    port: 10000,
  });
  console.log(result);
}
