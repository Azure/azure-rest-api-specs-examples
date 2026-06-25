const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a sandbox custom image.
 *
 * @summary creates or updates a sandbox custom image.
 * x-ms-original-file: 2025-02-14/KustoSandboxCustomImagesCreateOrUpdate.json
 */
async function kustoSandboxCustomImagesCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.sandboxCustomImages.createOrUpdate(
    "kustorptest",
    "kustoCluster",
    "customImage8",
    { languageVersion: "3.10.8", requirementsFileContent: "Requests", language: "Python" },
  );
  console.log(result);
}
