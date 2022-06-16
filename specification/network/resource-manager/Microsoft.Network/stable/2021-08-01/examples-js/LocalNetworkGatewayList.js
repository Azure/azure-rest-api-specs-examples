const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the local network gateways in a resource group.
 *
 * @summary Gets all the local network gateways in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LocalNetworkGatewayList.json
 */
async function listLocalNetworkGateways() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.localNetworkGateways.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listLocalNetworkGateways().catch(console.error);
