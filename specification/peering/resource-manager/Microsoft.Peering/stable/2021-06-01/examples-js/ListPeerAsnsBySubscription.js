const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the peer ASNs under the given subscription.
 *
 * @summary Lists all of the peer ASNs under the given subscription.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/ListPeerAsnsBySubscription.json
 */
async function listPeerAsNsInASubscription() {
  const subscriptionId = "subId";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peerAsns.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPeerAsNsInASubscription().catch(console.error);
