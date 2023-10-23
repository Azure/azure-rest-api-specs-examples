const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a sandbox custom image.
 *
 * @summary Updates a sandbox custom image.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImageUpdate.json
 */
async function kustoSandboxCustomImagesUpdate() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const sandboxCustomImageName = "customImage8";
  const parameters = {
    languageVersion: "3.10.8",
    requirementsFileContent: "Requests",
    language: "Python",
  };
  const credential = new DefaultAzureCredential();
  const client = new KustoManagementClient(credential, subscriptionId);
  const result = await client.sandboxCustomImages.beginUpdateAndWait(
    resourceGroupName,
    clusterName,
    sandboxCustomImageName,
    parameters
  );
  console.log(result);
}
