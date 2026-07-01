const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resource group level policy assignment.
 *
 * @summary queries policy events for the resource group level policy assignment.
 * x-ms-original-file: 2024-10-01/PolicyEvents_QueryResourceGroupLevelPolicyAssignmentScope.json
 */
async function queryAtResourceGroupLevelPolicyAssignmentScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForResourceGroupLevelPolicyAssignment(
    "default",
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
    "myResourceGroup",
    "myPolicyAssignment",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
