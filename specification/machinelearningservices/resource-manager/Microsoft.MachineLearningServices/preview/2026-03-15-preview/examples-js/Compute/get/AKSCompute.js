const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets compute definition by its name. Any secrets (storage keys, service credentials, etc) are not returned - use 'keys' nested resource to get them.
 *
 * @summary gets compute definition by its name. Any secrets (storage keys, service credentials, etc) are not returned - use 'keys' nested resource to get them.
 * x-ms-original-file: 2026-03-15-preview/Compute/get/AKSCompute.json
 */
async function getAAKSCompute() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.computeOperations.get("testrg123", "workspaces123", "compute123");
  console.log(result);
}
