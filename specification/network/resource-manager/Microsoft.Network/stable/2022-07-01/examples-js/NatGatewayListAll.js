const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Nat Gateways in a subscription.
 *
 * @summary Gets all the Nat Gateways in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/NatGatewayListAll.json
 */
async function listAllNatGateways() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.natGateways.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllNatGateways().catch(console.error);
