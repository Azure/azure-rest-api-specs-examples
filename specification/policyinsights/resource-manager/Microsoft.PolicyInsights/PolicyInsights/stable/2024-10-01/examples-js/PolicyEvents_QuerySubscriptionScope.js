const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resources under the subscription.
 *
 * @summary queries policy events for the resources under the subscription.
 * x-ms-original-file: 2024-10-01/PolicyEvents_QuerySubscriptionScope.json
 */
async function queryAtSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForSubscription(
    "default",
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
