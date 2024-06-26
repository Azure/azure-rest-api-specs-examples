const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the VpnServerConfigurations in a subscription.
 *
 * @summary Lists all the VpnServerConfigurations in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VpnServerConfigurationList.json
 */
async function vpnServerConfigurationList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.vpnServerConfigurations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

vpnServerConfigurationList().catch(console.error);
