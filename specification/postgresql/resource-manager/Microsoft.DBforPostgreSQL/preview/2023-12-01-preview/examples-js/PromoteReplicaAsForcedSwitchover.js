const { PostgreSQLManagementFlexibleServerClient } = require("@azure/arm-postgresql-flexible");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 *
 * @summary Updates an existing server. The request body can contain one to many of the properties present in the normal server definition.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-12-01-preview/examples/PromoteReplicaAsForcedSwitchover.json
 */
async function switchOverAReplicaServerAsForcedIEItWillReplicaAsPrimaryAndOriginalPrimaryAsReplicaImmediatelyWithoutWaitingForPrimaryAndReplicaToBeInSync() {
  const subscriptionId =
    process.env["POSTGRESQL_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["POSTGRESQL_RESOURCE_GROUP"] || "testResourceGroup";
  const serverName = "pgtestsvc4-replica";
  const parameters = {
    replica: { promoteMode: "switchover", promoteOption: "forced" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementFlexibleServerClient(credential, subscriptionId);
  const result = await client.servers.beginUpdateAndWait(resourceGroupName, serverName, parameters);
  console.log(result);
}
