const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all remediations for the management group.
 *
 * @summary Gets all remediations for the management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2021-10-01/examples/Remediations_ListManagementGroupScope.json
 */
async function listRemediationsAtManagementGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "financeMg";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.remediations.listForManagementGroup(managementGroupId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRemediationsAtManagementGroupScope().catch(console.error);
