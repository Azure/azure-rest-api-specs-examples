const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Copies an image to this container registry from the specified container registry.
 *
 * @summary Copies an image to this container registry from the specified container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/ImportImageByTag.json
 */
async function importImageByTag() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const parameters = {
    mode: "Force",
    source: {
      resourceId:
        "/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/sourceResourceGroup/providers/Microsoft.ContainerRegistry/registries/sourceRegistry",
      sourceImage: "sourceRepository:sourceTag",
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
