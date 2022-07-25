const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the container registry name is available for use. The name must contain only alphanumeric characters, be globally unique, and between 5 and 50 characters in length.
 *
 * @summary Checks whether the container registry name is available for use. The name must contain only alphanumeric characters, be globally unique, and between 5 and 50 characters in length.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/RegistryCheckNameNotAvailable.json
 */
async function registryCheckNameNotAvailable() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const registryNameCheckRequest = {
    name: "myRegistry",
    type: "Microsoft.ContainerRegistry/registries",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.checkNameAvailability(registryNameCheckRequest);
  console.log(result);
}

registryCheckNameNotAvailable().catch(console.error);
