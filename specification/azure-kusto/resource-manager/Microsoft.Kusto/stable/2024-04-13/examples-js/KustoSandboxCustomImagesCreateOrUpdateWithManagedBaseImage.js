const { KustoManagementClient } = require("@azure/arm-kusto");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a sandbox custom image.
 *
 * @summary Creates or updates a sandbox custom image.
 * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoSandboxCustomImagesCreateOrUpdateWithManagedBaseImage.json
 */
async function kustoSandboxCustomImagesCreateOrUpdateWithManagedBaseImage() {
  const subscriptionId =
    process.env["KUSTO_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = process.env["KUSTO_RESOURCE_GROUP"] || "kustorptest";
  const clusterName = "kustoCluster";
  const sandboxCustomImageName = "customImage2";
  const parameters = {
    baseImageName: "Python3_10_8",
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
