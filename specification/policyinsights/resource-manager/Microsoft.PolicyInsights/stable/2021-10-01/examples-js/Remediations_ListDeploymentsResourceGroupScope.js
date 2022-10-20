const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all deployments for a remediation at resource group scope.
 *
 * @summary Gets all deployments for a remediation at resource group scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListDeploymentsResourceGroupScope.json
 */
async function listDeploymentsForARemediationAtResourceGroupScope() {
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const resourceGroupName = "myResourceGroup";
  const remediationName = "myRemediation";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.remediations.listDeploymentsAtResourceGroup(
    resourceGroupName,
    remediationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDeploymentsForARemediationAtResourceGroupScope().catch(console.error);
