const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stops compute on all cluster nodes.
 *
 * @summary Stops compute on all cluster nodes.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterStop.json
 */
async function stopAllServersInTheCluster() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestGroup";
  const clusterName = "testcluster1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.beginStopAndWait(resourceGroupName, clusterName);
  console.log(result);
}
