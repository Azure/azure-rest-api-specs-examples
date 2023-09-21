const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a pipeline run for a container registry with the specified parameters
 *
 * @summary Creates a pipeline run for a container registry with the specified parameters
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-08-01-preview/examples/PipelineRunCreate_Export.json
 */
async function pipelineRunCreateExport() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const pipelineRunName = "myPipelineRun";
  const pipelineRunCreateParameters = {
    request: {
      artifacts: [
        "sourceRepository/hello-world",
        "sourceRepository2@sha256:00000000000000000000000000000000000",
      ],
      pipelineResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/exportPipelines/myExportPipeline",
      target: { name: "myblob.tar.gz", type: "AzureStorageBlob" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.pipelineRuns.beginCreateAndWait(
    resourceGroupName,
    registryName,
    pipelineRunName,
    pipelineRunCreateParameters
  );
  console.log(result);
}
