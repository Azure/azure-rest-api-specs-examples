const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new cluster with servers.
 *
 * @summary creates a new cluster with servers.
 * x-ms-original-file: 2023-03-02-preview/ClusterCreatePITR.json
 */
async function createANewClusterAsAPointInTimeRestore() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.create("TestGroup", "testcluster", {
    location: "westus",
    pointInTimeUTC: new Date("2017-12-14T00:00:37.467Z"),
    sourceLocation: "westus",
    sourceResourceId:
      "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/source-cluster",
  });
  console.log(result);
}
