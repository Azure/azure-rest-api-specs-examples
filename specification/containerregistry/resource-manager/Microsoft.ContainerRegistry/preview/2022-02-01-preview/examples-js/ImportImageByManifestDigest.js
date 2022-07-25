const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Copies an image to this container registry from the specified container registry.
 *
 * @summary Copies an image to this container registry from the specified container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ImportImageByManifestDigest.json
 */
async function importImageByManifestDigest() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const parameters = {
    mode: "Force",
    source: {
      resourceId:
        "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/sourceResourceGroup/providers/Microsoft.ContainerRegistry/registries/sourceRegistry",
      sourceImage:
        "sourceRepository@sha256:0000000000000000000000000000000000000000000000000000000000000000",
    },
    targetTags: ["targetRepository:targetTag"],
    untaggedTargetRepositories: ["targetRepository1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.beginImportImageAndWait(
    resourceGroupName,
    registryName,
    parameters
  );
  console.log(result);
}

importImageByManifestDigest().catch(console.error);
