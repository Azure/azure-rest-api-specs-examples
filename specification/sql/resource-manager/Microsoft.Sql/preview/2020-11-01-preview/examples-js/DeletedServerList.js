const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of deleted servers for a location.
 *
 * @summary Gets a list of deleted servers for a location.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DeletedServerList.json
 */
async function listDeletedServers() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const locationName = "japaneast";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedServers.listByLocation(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDeletedServers().catch(console.error);
