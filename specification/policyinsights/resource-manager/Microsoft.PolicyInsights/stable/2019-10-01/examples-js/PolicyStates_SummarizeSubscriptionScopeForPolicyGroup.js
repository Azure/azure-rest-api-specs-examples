const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Summarizes policy states for the resources under the subscription.
 *
 * @summary Summarizes policy states for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionScopeForPolicyGroup.json
 */
async function summarizeAtSubscriptionScopeForAPolicyDefinitionGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesSummaryResource = "latest";
  const top = 1;
  const filter = "'group1' IN PolicyDefinitionGroupNames";
  const options = {
    queryOptions: { top: top, filter: filter },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.summarizeForSubscription(
    policyStatesSummaryResource,
    subscriptionId,
    options
  );
  console.log(result);
}

summarizeAtSubscriptionScopeForAPolicyDefinitionGroup().catch(console.error);
