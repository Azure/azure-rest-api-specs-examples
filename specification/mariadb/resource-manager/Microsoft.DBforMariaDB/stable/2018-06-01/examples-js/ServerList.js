const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the servers in a given subscription.
 *
 * @summary List all the servers in a given subscription.
 * x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/ServerList.json
 */
async function serverList() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.servers.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

serverList().catch(console.error);
