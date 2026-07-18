const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to copies an image to this container registry from the specified container registry.
 *
 * @summary copies an image to this container registry from the specified container registry.
 * x-ms-original-file: 2026-03-01-preview/ImportImageFromPublicRegistry.json
 */
async function importImageFromPublicRegistry() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  await client.registries.importImage("myResourceGroup", "myRegistry", {
    source: { registryUri: "registry.hub.docker.com", sourceImage: "library/hello-world" },
    targetTags: ["targetRepository:targetTag"],
    untaggedTargetRepositories: ["targetRepository1"],
    mode: "Force",
  });
}
