const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries component policy states under subscription scope.
 *
 * @summary queries component policy states under subscription scope.
 * x-ms-original-file: 2024-10-01/ComponentPolicyStates_QuerySubscriptionScope.json
 */
async function queryLatestComponentPolicyStatesAtSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.componentPolicyStates.listQueryResultsForSubscription(
    "fff10b27-fff3-fff5-fff8-fffbe01e86a5",
    "latest",
  );
  console.log(result);
}
