const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the pipeline runs for the specified container registry.
 *
 * @summary Lists all the pipeline runs for the specified container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/PipelineRunList.json
 */
async function pipelineRunList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.pipelineRuns.list(resourceGroupName, registryName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

pipelineRunList().catch(console.error);
