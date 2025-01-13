const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes an existing remediation at management group scope.
 *
 * @summary Deletes an existing remediation at management group scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/Remediations_DeleteManagementGroupScope.json
 */
async function deleteRemediationAtManagementGroupScope() {
  const managementGroupId = "financeMg";
  const remediationName = "storageRemediation";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.remediations.deleteAtManagementGroup(
    managementGroupId,
    remediationName,
  );
  console.log(result);
}
