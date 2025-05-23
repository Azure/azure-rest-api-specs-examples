const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 *
 * @summary Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/PromoteReplicaAsPlannedSwitchover.json
 */
async function switchOverAReplicaServerAsPlannedIEItWillWaitForReplicationToCompleteBeforePromotingReplicaAsPrimaryAndOriginalPrimaryAsReplica() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testResourceGroup";
  const serverName = "pgtestsvc4-replica";
  const parameters = {
    replica: { promoteMode: "switchover", promoteOption: "planned" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginUpdateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
