const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates an export pipeline for a container registry with the specified parameters.
 *
 * @summary Creates an export pipeline for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ExportPipelineCreate.json
 */
async function exportPipelineCreate() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const exportPipelineName = "myExportPipeline";
  const exportPipelineCreateParameters = {
    identity: { type: "SystemAssigned" },
    location: "westus",
    options: ["OverwriteBlobs"],
    target: {
      type: "AzureStorageBlobContainer",
      keyVaultUri: "https://myvault.vault.azure.net/secrets/acrexportsas",
      uri: "https://accountname.blob.core.windows.net/containername",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.exportPipelines.beginCreateAndWait(
    resourceGroupName,
    registryName,
    exportPipelineName,
    exportPipelineCreateParameters,
  );
  console.log(result);
}
