const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy tracked resources under the subscription.
 *
 * @summary queries policy tracked resources under the subscription.
 * x-ms-original-file: 2018-07-01-preview/PolicyTrackedResources_QuerySubscriptionScope.json
 */
async function queryAtSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.policyTrackedResources.listQueryResultsForSubscription(
    "default",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
