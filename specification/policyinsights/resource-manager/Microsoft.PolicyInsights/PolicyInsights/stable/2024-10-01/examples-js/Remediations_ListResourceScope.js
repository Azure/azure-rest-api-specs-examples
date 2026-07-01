const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets all remediations for a resource.
 *
 * @summary gets all remediations for a resource.
 * x-ms-original-file: 2024-10-01/Remediations_ListResourceScope.json
 */
async function listRemediationsAtIndividualResourceScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.remediations.listForResource(
    "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myResourceGroup/providers/microsoft.storage/storageaccounts/storAc1",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
