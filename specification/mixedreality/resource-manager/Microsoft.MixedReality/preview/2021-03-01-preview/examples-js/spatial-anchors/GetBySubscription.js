const { MixedRealityClient } = require("@azure/arm-mixedreality");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Spatial Anchors Accounts by Subscription
 *
 * @summary List Spatial Anchors Accounts by Subscription
 * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/spatial-anchors/GetBySubscription.json
 */
async function listSpatialAnchorsAccountsBySubscription() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const credential = new DefaultAzureCredential();
  const client = new MixedRealityClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.spatialAnchorsAccounts.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSpatialAnchorsAccountsBySubscription().catch(console.error);
