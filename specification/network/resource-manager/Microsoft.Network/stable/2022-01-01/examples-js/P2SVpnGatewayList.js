const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the P2SVpnGateways in a subscription.
 *
 * @summary Lists all the P2SVpnGateways in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/P2SVpnGatewayList.json
 */
async function p2SVpnGatewayListBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.p2SVpnGateways.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

p2SVpnGatewayListBySubscription().catch(console.error);
