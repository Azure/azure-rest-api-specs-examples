const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Recovers a deleted server.
 *
 * @summary Recovers a deleted server.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerRecover.json
 */
async function recoverDeletedServer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "japaneast";
  const deletedServerName = "sqlcrudtest-d-1414";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.deletedServers.beginRecoverAndWait(locationName, deletedServerName);
  console.log(result);
}

recoverDeletedServer().catch(console.error);
