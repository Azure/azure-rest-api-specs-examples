const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of AvailabilitySets in a subscription.
 *
 * @summary List of AvailabilitySets in a subscription.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/AvailabilitySets_ListBySubscription_MinimumSet_Gen.json
 */
async function availabilitySetsListBySubscriptionMinimumSet() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "79332E5A-630B-480F-A266-A941C015AB19";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilitySets.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
