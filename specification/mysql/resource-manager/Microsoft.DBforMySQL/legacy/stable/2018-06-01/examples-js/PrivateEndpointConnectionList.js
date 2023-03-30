const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all private endpoint connections on a server.
 *
 * @summary Gets all private endpoint connections on a server.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateEndpointConnectionList.json
 */
async function getsListOfPrivateEndpointConnectionsOnAServer() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const serverName = "test-svr";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByServer(
    resourceGroupName,
    serverName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsListOfPrivateEndpointConnectionsOnAServer().catch(console.error);
