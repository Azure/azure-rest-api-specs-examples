const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Object Anchors Accounts by Subscription
 *
 * @summary List Object Anchors Accounts by Subscription
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/object-anchors/GetBySubscription.json
 */
async function listObjectAnchorsAccountsBySubscription() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.objectAnchorsAccounts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listObjectAnchorsAccountsBySubscription().catch(console.error);
