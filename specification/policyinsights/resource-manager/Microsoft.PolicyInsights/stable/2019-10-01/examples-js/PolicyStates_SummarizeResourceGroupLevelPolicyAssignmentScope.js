const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Summarizes policy states for the resource group level policy assignment.
 *
 * @summary Summarizes policy states for the resource group level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeResourceGroupLevelPolicyAssignmentScope.json
 */
async function summarizeAtPolicyAssignmentScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyStatesSummaryResource = "latest";
  const resourceGroupName = process.env["POLICYINSIGHTS_RESOURCE_GROUP"] || "myResourceGroup";
  const policyAssignmentName = "b7a1ca2596524e3ab19597f2";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.summarizeForResourceGroupLevelPolicyAssignment(
    policyStatesSummaryResource,
    subscriptionId,
    resourceGroupName,
    policyAssignmentName
  );
  console.log(result);
}
