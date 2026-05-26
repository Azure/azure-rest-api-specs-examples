const { RedisEnterpriseManagementClient } = require("@azure/arm-redisenterprisecache");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a database
 *
 * @summary creates a database
 * x-ms-original-file: 2026-02-01-preview/RedisEnterpriseDatabasesNoClusterCacheCreate.json
 */
async function redisEnterpriseDatabasesCreateNoClusterCache() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "e7b5a9d2-6b6a-4d2f-9143-20d9a10f5b8f";
  const client = new RedisEnterpriseManagementClient(credential, subscriptionId);
  const result = await client.databases.create("rg1", "cache1", "default", {
    clientProtocol: "Encrypted",
    clusteringPolicy: "NoCluster",
    evictionPolicy: "NoEviction",
    notifyKeyspaceEvents: "",
    port: 10000,
  });
  console.log(result);
}
