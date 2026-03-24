const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to copies an image to this container registry from the specified container registry.
 *
 * @summary copies an image to this container registry from the specified container registry.
 * x-ms-original-file: 2026-01-01-preview/ImportImageByTag.json
 */
async function importImageByTag() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  await client.registries.importImage("myResourceGroup", "myRegistry", {
    source: {
      resourceId:
        "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/sourceResourceGroup/providers/Microsoft.ContainerRegistry/registries/sourceRegistry",
      sourceImage: "sourceRepository:sourceTag",
    },
    targetTags: ["targetRepository:targetTag"],
    untaggedTargetRepositories: ["targetRepository1"],
    mode: "Force",
  });
}
