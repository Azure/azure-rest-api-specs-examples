const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the detailed information for a given pipeline run.
 *
 * @summary Gets the detailed information for a given pipeline run.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/PipelineRunGet.json
 */
async function pipelineRunGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const pipelineRunName = "myPipelineRun";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.pipelineRuns.get(resourceGroupName, registryName, pipelineRunName);
  console.log(result);
}

pipelineRunGet().catch(console.error);
