const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the upload location for the user to be able to upload the source.
 *
 * @summary Get the upload location for the user to be able to upload the source.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesGetBuildSourceUploadUrl.json
 */
async function registriesGetBuildSourceUploadUrl() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.getBuildSourceUploadUrl(resourceGroupName, registryName);
  console.log(result);
}

registriesGetBuildSourceUploadUrl().catch(console.error);
