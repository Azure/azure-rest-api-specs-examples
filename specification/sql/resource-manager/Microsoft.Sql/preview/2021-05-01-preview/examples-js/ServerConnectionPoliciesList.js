const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists connection policy
 *
 * @summary Lists connection policy
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ServerConnectionPoliciesList.json
 */
async function listsAServersConnectionPolicies() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "rgtest-12";
  const serverName = "servertest-6285";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverConnectionPolicies.listByServer(
    resourceGroupName,
    serverName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAServersConnectionPolicies().catch(console.error);
