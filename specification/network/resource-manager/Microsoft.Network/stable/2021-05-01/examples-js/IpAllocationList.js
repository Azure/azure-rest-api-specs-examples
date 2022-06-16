const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all IpAllocations in a subscription.
 *
 * @summary Gets all IpAllocations in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/IpAllocationList.json
 */
async function listAllIPAllocations() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.ipAllocations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllIPAllocations().catch(console.error);
