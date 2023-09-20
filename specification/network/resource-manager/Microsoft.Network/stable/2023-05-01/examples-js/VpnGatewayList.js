const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the VpnGateways in a subscription.
 *
 * @summary Lists all the VpnGateways in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/VpnGatewayList.json
 */
async function vpnGatewayListBySubscription() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vpnGateways.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
