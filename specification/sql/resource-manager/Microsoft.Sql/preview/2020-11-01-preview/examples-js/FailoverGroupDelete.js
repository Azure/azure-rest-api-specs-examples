const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a failover group.
 *
 * @summary Deletes a failover group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/FailoverGroupDelete.json
 */
async function deleteFailoverGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "failover-group-primary-server";
  const failoverGroupName = "failover-group-test-1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.failoverGroups.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    failoverGroupName
  );
  console.log(result);
}

deleteFailoverGroup().catch(console.error);
