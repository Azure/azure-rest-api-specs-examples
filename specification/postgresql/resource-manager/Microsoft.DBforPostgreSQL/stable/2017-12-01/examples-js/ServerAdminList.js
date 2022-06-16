const { PostgreSQLManagementClient } = require("@azure/arm-postgresql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of server Administrators.
 *
 * @summary Returns a list of server Administrators.
 * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerAdminList.json
 */
async function getAListOfServerAdministrators() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "pgtestsvc4";
  const credential = new DefaultAzureCredential();
  const client = new PostgreSQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverAdministrators.list(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAListOfServerAdministrators().catch(console.error);
