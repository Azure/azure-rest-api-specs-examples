const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the peerings under the given subscription.
 *
 * @summary Lists all of the peerings under the given subscription.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListPeeringsBySubscription.json
 */
async function listPeeringsInASubscription() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peerings.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPeeringsInASubscription().catch(console.error);
