const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets all attestations for the resource group.
 *
 * @summary gets all attestations for the resource group.
 * x-ms-original-file: 2024-10-01/Attestations_ListResourceGroupScope.json
 */
async function listAttestationsAtResourceGroupScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.attestations.listForResourceGroup("myRg")) {
    resArray.push(item);
  }

  console.log(resArray);
}
