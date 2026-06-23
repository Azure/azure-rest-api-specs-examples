const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new cluster with servers.
 *
 * @summary creates a new cluster with servers.
 * x-ms-original-file: 2023-03-02-preview/ClusterCreateMultiNode.json
 */
async function createANewMultiNodeCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.create("TestGroup", "testcluster-multinode", {
    location: "westus",
    administratorLoginPassword: "password",
    citusVersion: "11.1",
    coordinatorEnablePublicIpAccess: true,
    coordinatorServerEdition: "GeneralPurpose",
    coordinatorStorageQuotaInMb: 524288,
    coordinatorVCores: 4,
    enableHa: true,
    enableShardsOnCoordinator: false,
    nodeCount: 3,
    nodeEnablePublicIpAccess: false,
    nodeServerEdition: "MemoryOptimized",
    nodeStorageQuotaInMb: 524288,
    nodeVCores: 8,
    postgresqlVersion: "15",
    preferredPrimaryZone: "1",
    tags: {},
  });
  console.log(result);
}
