const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers a policy evaluation scan for all the resources under the subscription
 *
 * @summary Triggers a policy evaluation scan for all the resources under the subscription
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_TriggerSubscriptionEvaluation.json
 */
async function triggerEvaluationsForAllResourcesInASubscription() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.beginTriggerSubscriptionEvaluationAndWait(
    subscriptionId
  );
  console.log(result);
}
