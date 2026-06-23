const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing cluster. The request body can contain one or several properties from the cluster definition.
 *
 * @summary updates an existing cluster. The request body can contain one or several properties from the cluster definition.
 * x-ms-original-file: 2023-03-02-preview/ClusterScaleStorage.json
 */
async function scaleUpStorage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.update("TestGroup", "testcluster", {
    nodeStorageQuotaInMb: 2097152,
  });
  console.log(result);
}
