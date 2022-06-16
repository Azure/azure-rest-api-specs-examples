const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all Network Virtual Appliances in a subscription.
 *
 * @summary Gets all Network Virtual Appliances in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkVirtualApplianceListBySubscription.json
 */
async function listAllNetworkVirtualAppliancesForAGivenSubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkVirtualAppliances.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllNetworkVirtualAppliancesForAGivenSubscription().catch(console.error);
