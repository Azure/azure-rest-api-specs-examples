const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of server Administrators.
 *
 * @summary Returns a list of server Administrators.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2017-12-01/examples/ServerAdminList.json
 */
async function getAListOfServerAdministrators() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testrg";
  const serverName = "mysqltestsvc4";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serverAdministrators.list(resourceGroupName, serverName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAListOfServerAdministrators().catch(console.error);
