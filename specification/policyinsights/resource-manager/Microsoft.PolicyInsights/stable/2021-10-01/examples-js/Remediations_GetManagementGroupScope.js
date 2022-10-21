const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing remediation at management group scope.
 *
 * @summary Gets an existing remediation at management group scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_GetManagementGroupScope.json
 */
async function getRemediationAtManagementGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "financeMg";
  const remediationName = "storageRemediation";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.remediations.getAtManagementGroup(managementGroupId, remediationName);
  console.log(result);
}

getRemediationAtManagementGroupScope().catch(console.error);
