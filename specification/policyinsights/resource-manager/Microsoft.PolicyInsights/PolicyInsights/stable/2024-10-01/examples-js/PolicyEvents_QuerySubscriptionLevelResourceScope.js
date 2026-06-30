const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resource.
 *
 * @summary queries policy events for the resource.
 * x-ms-original-file: 2024-10-01/PolicyEvents_QuerySubscriptionLevelResourceScope.json
 */
async function queryAtSubscriptionLevelResourceScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForResource(
    "default",
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.SomeNamespace/someResourceType/someResourceName",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
