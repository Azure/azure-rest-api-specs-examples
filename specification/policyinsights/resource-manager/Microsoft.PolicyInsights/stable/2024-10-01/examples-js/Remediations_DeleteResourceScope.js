const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes an existing remediation at individual resource scope.
 *
 * @summary Deletes an existing remediation at individual resource scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/Remediations_DeleteResourceScope.json
 */
async function deleteRemediationAtIndividualResourceScope() {
  const resourceId =
    "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1";
  const remediationName = "storageRemediation";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.remediations.deleteAtResource(resourceId, remediationName);
  console.log(result);
}
