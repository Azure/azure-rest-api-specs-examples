const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a new cluster with servers.
 *
 * @summary creates a new cluster with servers.
 * x-ms-original-file: 2023-03-02-preview/ClusterCreateBurstablev2.json
 */
async function createANewSingleNodeBurstable2VCoresCluster() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.create("TestGroup", "testcluster-burstablev2", {
    location: "westus",
    administratorLoginPassword: "password",
    citusVersion: "11.3",
    coordinatorEnablePublicIpAccess: true,
    coordinatorServerEdition: "BurstableGeneralPurpose",
    coordinatorStorageQuotaInMb: 131072,
    coordinatorVCores: 2,
    enableHa: false,
    enableShardsOnCoordinator: true,
    nodeCount: 0,
    postgresqlVersion: "15",
    preferredPrimaryZone: "1",
    tags: { owner: "JohnDoe" },
  });
  console.log(result);
}
