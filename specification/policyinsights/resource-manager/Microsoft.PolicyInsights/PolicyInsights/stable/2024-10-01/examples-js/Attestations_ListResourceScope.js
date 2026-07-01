const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets all attestations for a resource.
 *
 * @summary gets all attestations for a resource.
 * x-ms-original-file: 2024-10-01/Attestations_ListResourceScope.json
 */
async function listAttestationsAtIndividualResourceScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.attestations.listForResource(
    "subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/resourcegroups/myrg/providers/microsoft.compute/virtualMachines/devVM",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
