const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate keys for a token of a specified container registry.
 *
 * @summary Generate keys for a token of a specified container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/RegistryGenerateCredentials.json
 */
async function registryGenerateCredentials() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const generateCredentialsParameters = {
    expiry: new Date("2020-12-31T15:59:59.0707808Z"),
    tokenId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/myToken",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.beginGenerateCredentialsAndWait(
    resourceGroupName,
    registryName,
    generateCredentialsParameters
  );
  console.log(result);
}

registryGenerateCredentials().catch(console.error);
