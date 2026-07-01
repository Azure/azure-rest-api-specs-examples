const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates the specified managed compute deployment associated with the Cognitive Services account.
 *
 * @summary updates the specified managed compute deployment associated with the Cognitive Services account.
 * x-ms-original-file: 2026-05-15-preview/UpdateManagedComputeDeployment.json
 */
async function updateManagedComputeDeployment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.managedComputeDeployments.update(
    "resourceGroupName",
    "accountName",
    "gpt-oss-120b-gpu",
    { sku: { name: "GlobalManagedCompute", capacity: 2 } },
  );
  console.log(result);
}
