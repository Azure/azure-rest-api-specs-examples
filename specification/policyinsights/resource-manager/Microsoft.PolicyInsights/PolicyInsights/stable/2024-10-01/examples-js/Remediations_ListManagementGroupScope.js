const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets all remediations for the management group.
 *
 * @summary gets all remediations for the management group.
 * x-ms-original-file: 2024-10-01/Remediations_ListManagementGroupScope.json
 */
async function listRemediationsAtManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.remediations.listForManagementGroup("financeMg")) {
    resArray.push(item);
  }

  console.log(resArray);
}
