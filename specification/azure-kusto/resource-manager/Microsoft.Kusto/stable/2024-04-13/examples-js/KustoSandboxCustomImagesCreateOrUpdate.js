const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a sandbox custom image.
 *
 * @summary Creates or updates a sandbox custom image.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoSandboxCustomImagesCreateOrUpdate.json
 */
async function kustoSandboxCustomImagesCreateOrUpdate() {
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
  const result = await client.sandboxCustomImages.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    sandboxCustomImageName,
    parameters,
  );
  console.log(result);
}
