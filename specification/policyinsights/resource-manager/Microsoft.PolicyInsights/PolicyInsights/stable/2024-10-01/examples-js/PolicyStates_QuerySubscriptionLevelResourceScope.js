const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy states for the resource.
 *
 * @summary queries policy states for the resource.
 * x-ms-original-file: 2024-10-01/PolicyStates_QuerySubscriptionLevelResourceScope.json
 */
async function queryAllPolicyStatesAtSubscriptionLevelResourceScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyStates.listQueryResultsForResource(
    "default",
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.SomeNamespace/someResourceType/someResourceName",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
