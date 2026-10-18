const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to drains requests and deletes an adapter deployment and its serving snapshot.
 *
 * @summary drains requests and deletes an adapter deployment and its serving snapshot.
 * x-ms-original-file: 2026-09-15-preview/DeleteAdapterDeployment.json
 */
async function deleteAdapterDeployment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  await client.adapterDeployments.delete(
    "resourceGroupName",
    "accountName",
    "adapterDeploymentName",
    { ifMatch: '"0x8D1234567890ABC"' },
  );
}
