const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates a pipeline run for a container registry with the specified parameters
 *
 * @summary Creates a pipeline run for a container registry with the specified parameters
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/PipelineRunCreate_Import.json
 */
async function pipelineRunCreateImport() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const pipelineRunName = "myPipelineRun";
  const pipelineRunCreateParameters = {
    forceUpdateTag: "2020-03-04T17:23:21.9261521+00:00",
    request: {
      catalogDigest: "sha256@",
      pipelineResourceId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline",
      source: { name: "myblob.tar.gz", type: "AzureStorageBlob" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.pipelineRuns.beginCreateAndWait(
    resourceGroupName,
    registryName,
    pipelineRunName,
    pipelineRunCreateParameters,
  );
  console.log(result);
}
