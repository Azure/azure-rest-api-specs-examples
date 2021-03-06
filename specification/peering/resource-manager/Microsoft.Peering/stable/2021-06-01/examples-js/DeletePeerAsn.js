const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing peer ASN with the specified name under the given subscription.
 *
 * @summary Deletes an existing peer ASN with the specified name under the given subscription.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/DeletePeerAsn.json
 */
async function deleteAPeerAsn() {
  const subscriptionId = "subId";
  const peerAsnName = "peerAsnName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerAsns.delete(peerAsnName);
  console.log(result);
}

deleteAPeerAsn().catch(console.error);
