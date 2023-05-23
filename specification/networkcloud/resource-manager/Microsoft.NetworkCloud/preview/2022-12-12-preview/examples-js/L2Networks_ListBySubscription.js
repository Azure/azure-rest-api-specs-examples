const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of layer 2 (L2) networks in the provided subscription.
 *
 * @summary Get a list of layer 2 (L2) networks in the provided subscription.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/L2Networks_ListBySubscription.json
 */
async function listL2NetworksForSubscription() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.l2Networks.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
