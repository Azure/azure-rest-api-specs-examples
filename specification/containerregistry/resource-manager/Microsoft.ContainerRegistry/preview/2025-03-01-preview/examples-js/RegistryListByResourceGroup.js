const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all the container registries under the specified resource group.
 *
 * @summary Lists all the container registries under the specified resource group.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/RegistryListByResourceGroup.json
 */
async function registryListByResourceGroup() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.registries.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
