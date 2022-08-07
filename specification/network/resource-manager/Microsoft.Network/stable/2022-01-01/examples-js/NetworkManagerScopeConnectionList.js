const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all scope connections created by this network manager.
 *
 * @summary List all scope connections created by this network manager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerScopeConnectionList.json
 */
async function listNetworkManagerScopeConnection() {
  const subscriptionId = "subscriptionC";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scopeConnections.list(resourceGroupName, networkManagerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listNetworkManagerScopeConnection().catch(console.error);
