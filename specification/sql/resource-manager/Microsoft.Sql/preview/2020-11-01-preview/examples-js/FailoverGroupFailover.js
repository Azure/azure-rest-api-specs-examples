const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Fails over from the current primary server to this server.
 *
 * @summary Fails over from the current primary server to this server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupFailover.json
 */
async function plannedFailoverOfAFailoverGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "failover-group-secondary-server";
  const failoverGroupName = "failover-group-test-3";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.failoverGroups.beginFailoverAndWait(
    resourceGroupName,
    serverName,
    failoverGroupName
  );
  console.log(result);
}

plannedFailoverOfAFailoverGroup().catch(console.error);
