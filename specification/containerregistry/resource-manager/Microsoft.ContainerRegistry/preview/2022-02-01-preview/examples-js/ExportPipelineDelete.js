const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an export pipeline from a container registry.
 *
 * @summary Deletes an export pipeline from a container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ExportPipelineDelete.json
 */
async function exportPipelineDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const exportPipelineName = "myExportPipeline";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.exportPipelines.beginDeleteAndWait(
    resourceGroupName,
    registryName,
    exportPipelineName
  );
  console.log(result);
}

exportPipelineDelete().catch(console.error);
