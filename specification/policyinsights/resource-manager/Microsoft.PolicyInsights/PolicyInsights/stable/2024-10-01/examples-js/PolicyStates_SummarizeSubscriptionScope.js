const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to summarizes policy states for the resources under the subscription.
 *
 * @summary summarizes policy states for the resources under the subscription.
 * x-ms-original-file: 2024-10-01/PolicyStates_SummarizeSubscriptionScope.json
 */
async function summarizeAtSubscriptionScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.policyStates.summarizeForSubscription(
    "latest",
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
    { queryOptions: { top: 5 } },
  );
  console.log(result);
}
