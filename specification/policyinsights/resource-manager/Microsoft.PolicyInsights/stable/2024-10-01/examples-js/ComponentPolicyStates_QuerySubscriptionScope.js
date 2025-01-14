const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries component policy states under subscription scope.
 *
 * @summary Queries component policy states under subscription scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QuerySubscriptionScope.json
 */
async function queryLatestComponentPolicyStatesAtSubscriptionScope() {
  const subscriptionId = "fff10b27-fff3-fff5-fff8-fffbe01e86a5";
  const componentPolicyStatesResource = "latest";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.componentPolicyStates.listQueryResultsForSubscription(
    subscriptionId,
    componentPolicyStatesResource,
  );
  console.log(result);
}
