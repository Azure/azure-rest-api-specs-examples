const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing server. The request body can contain one or multiple of the properties present in the normal server definition.
 *
 * @summary updates an existing server. The request body can contain one or multiple of the properties present in the normal server definition.
 * x-ms-original-file: 2026-01-01-preview/ServersPromoteReplicaAsPlannedSwitchover.json
 */
async function switchOverAReadReplicaToPrimaryServerWithPlannedDataSynchronizationMeaningThatItWaitsForDataInTheReadReplicaToBeFullySynchronizedWithItsSourceServerBeforeItInitiatesTheSwitchingOfRolesBetweenTheReadReplicaAndThePrimaryServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.update("exampleresourcegroup", "exampleserver", {
    replica: { promoteMode: "Switchover", promoteOption: "Planned" },
  });
  console.log(result);
}
