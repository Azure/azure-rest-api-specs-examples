const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a link to download the run logs.
 *
 * @summary Gets a link to download the run logs.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsGetLogSasUrl.json
 */
async function runsGetLogSasUrl() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const runId = "0accec26-d6de-4757-8e74-d080f38eaaab";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.runs.getLogSasUrl(resourceGroupName, registryName, runId);
  console.log(result);
}

runsGetLogSasUrl().catch(console.error);
