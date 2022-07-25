const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the runs for a registry.
 *
 * @summary Gets all the runs for a registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsList.json
 */
async function runsList() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const filter = "";
  const top = 10;
  const options = { filter, top };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.runs.list(resourceGroupName, registryName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

runsList().catch(console.error);
