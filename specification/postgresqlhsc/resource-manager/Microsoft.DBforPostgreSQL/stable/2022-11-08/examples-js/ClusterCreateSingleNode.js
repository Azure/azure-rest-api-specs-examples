const { CosmosDBForPostgreSQL } = require("@azure/arm-cosmosdbforpostgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new cluster with servers.
 *
 * @summary Creates a new cluster with servers.
 * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterCreateSingleNode.json
 */
async function createANewSingleNodeCluster() {
  const subscriptionId =
    process.env["COSMOSFORPOSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["COSMOSFORPOSTGRESQL_RESOURCE_GROUP"] || "TestGroup";
  const clusterName = "testcluster-singlenode";
  const parameters = {
    administratorLoginPassword: "password",
    citusVersion: "11.3",
    coordinatorEnablePublicIpAccess: true,
    coordinatorServerEdition: "GeneralPurpose",
    coordinatorStorageQuotaInMb: 131072,
    coordinatorVCores: 8,
    enableHa: true,
    enableShardsOnCoordinator: true,
    location: "westus",
    nodeCount: 0,
    postgresqlVersion: "15",
    preferredPrimaryZone: "1",
    tags: { owner: "JohnDoe" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBForPostgreSQL(credential, subscriptionId);
  const result = await client.clusters.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}
