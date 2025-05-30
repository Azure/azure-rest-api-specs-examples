const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the detailed information for a given task run.
 *
 * @summary Gets the detailed information for a given task run.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsGet.json
 */
async function taskRunsGet() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const taskRunName = "myRun";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.taskRuns.get(resourceGroupName, registryName, taskRunName);
  console.log(result);
}
