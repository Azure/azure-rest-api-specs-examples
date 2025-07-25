const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of layer 3 (L3) networks in the provided subscription.
 *
 * @summary Get a list of layer 3 (L3) networks in the provided subscription.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/L3Networks_ListBySubscription.json
 */
async function listL3NetworksForSubscription() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.l3Networks.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
