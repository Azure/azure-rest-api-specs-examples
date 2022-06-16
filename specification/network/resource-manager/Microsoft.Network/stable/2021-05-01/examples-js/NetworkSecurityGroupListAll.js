const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all network security groups in a subscription.
 *
 * @summary Gets all network security groups in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkSecurityGroupListAll.json
 */
async function listAllNetworkSecurityGroups() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkSecurityGroups.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllNetworkSecurityGroups().catch(console.error);
