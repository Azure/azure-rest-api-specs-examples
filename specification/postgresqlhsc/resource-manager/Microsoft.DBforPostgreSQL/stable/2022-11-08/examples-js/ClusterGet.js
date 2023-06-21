const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about a cluster such as compute and storage configuration and cluster lifecycle metadata such as cluster creation date and time.
 *
 * @summary Gets information about a cluster such as compute and storage configuration and cluster lifecycle metadata such as cluster creation date and time.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterGet.json
 */
async function getTheCluster() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestGroup";
  const clusterName = "testcluster1";
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.get(resourceGroupName, clusterName);
  console.log(result);
}
