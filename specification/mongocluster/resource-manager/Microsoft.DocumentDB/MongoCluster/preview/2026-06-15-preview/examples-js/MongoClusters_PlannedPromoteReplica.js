const { MongoClusterManagementClient } = require("@azure/arm-mongocluster");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to promotes a replica mongo cluster to a primary role.
 *
 * @summary promotes a replica mongo cluster to a primary role.
 * x-ms-original-file: 2026-06-15-preview/MongoClusters_PlannedPromoteReplica.json
 */
async function promotesAReplicaMongoClusterResourceToAPrimaryRoleWaitingForTheReplicaToCatchUpBeforePromoting() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new MongoClusterManagementClient(credential, subscriptionId);
  await client.mongoClusters.promote("TestGroup", "myMongoCluster", {
    promoteOption: "Planned",
    mode: "Switchover",
  });
}
