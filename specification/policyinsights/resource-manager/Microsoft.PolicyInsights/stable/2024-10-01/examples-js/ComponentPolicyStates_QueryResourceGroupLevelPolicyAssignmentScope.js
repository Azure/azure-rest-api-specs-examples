const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries component policy states for the resource group level policy assignment.
 *
 * @summary Queries component policy states for the resource group level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceGroupLevelPolicyAssignmentScope.json
 */
async function queryLatestAtResourceGroupLevelPolicyAssignmentScope() {
  const subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const resourceGroupName = process.env["POLICYINSIGHTS_RESOURCE_GROUP"] || "myResourceGroup";
  const policyAssignmentName = "myPolicyAssignment";
  const componentPolicyStatesResource = "latest";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result =
    await client.componentPolicyStates.listQueryResultsForResourceGroupLevelPolicyAssignment(
      subscriptionId,
      resourceGroupName,
      policyAssignmentName,
      componentPolicyStatesResource,
    );
  console.log(result);
}
