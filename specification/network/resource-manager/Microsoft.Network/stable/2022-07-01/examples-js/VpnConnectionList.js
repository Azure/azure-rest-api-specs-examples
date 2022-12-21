const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all vpn connections for a particular virtual wan vpn gateway.
 *
 * @summary Retrieves all vpn connections for a particular virtual wan vpn gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VpnConnectionList.json
 */
async function vpnConnectionList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vpnConnections.listByVpnGateway(resourceGroupName, gatewayName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

vpnConnectionList().catch(console.error);
