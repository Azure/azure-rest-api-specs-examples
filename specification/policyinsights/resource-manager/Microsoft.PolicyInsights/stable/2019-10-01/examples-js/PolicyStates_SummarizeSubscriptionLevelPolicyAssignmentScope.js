const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Summarizes policy states for the subscription level policy assignment.
 *
 * @summary Summarizes policy states for the subscription level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeSubscriptionLevelPolicyAssignmentScope.json
 */
async function summarizeAtPolicyAssignmentScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyStatesSummaryResource = "latest";
  const policyAssignmentName = "ec8f9645-8ecb-4abb-9c0b-5292f19d4003";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.summarizeForSubscriptionLevelPolicyAssignment(
    policyStatesSummaryResource,
    subscriptionId,
    policyAssignmentName
  );
  console.log(result);
}
