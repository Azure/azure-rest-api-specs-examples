const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries component policy states for the subscription level policy assignment.
 *
 * @summary Queries component policy states for the subscription level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/ComponentPolicyStates_QuerySubscriptionLevelPolicyAssignmentScope.json
 */
async function queryLatestAtSubscriptionLevelPolicyAssignmentScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyAssignmentName = "ec8f9645-8ecb-4abb-9c0b-5292f19d4003";
  const componentPolicyStatesResource = "latest";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result =
    await client.componentPolicyStates.listQueryResultsForSubscriptionLevelPolicyAssignment(
      subscriptionId,
      policyAssignmentName,
      componentPolicyStatesResource
    );
  console.log(result);
}
