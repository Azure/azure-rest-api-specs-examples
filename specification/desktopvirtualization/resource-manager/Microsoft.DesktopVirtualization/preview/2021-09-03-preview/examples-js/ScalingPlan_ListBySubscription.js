const { DesktopVirtualizationAPIClient } = require("@azure/arm-desktopvirtualization");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List scaling plans in subscription.
 *
 * @summary List scaling plans in subscription.
 * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ScalingPlan_ListBySubscription.json
 */
async function scalingPlansListBySubscription() {
  const subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
  const credential = new DefaultAzureCredential();
  const client = new DesktopVirtualizationAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.scalingPlans.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

scalingPlansListBySubscription().catch(console.error);
