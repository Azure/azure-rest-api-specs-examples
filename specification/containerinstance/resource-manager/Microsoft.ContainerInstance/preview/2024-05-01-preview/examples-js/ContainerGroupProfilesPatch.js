const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches container group profile with specified properties.
 *
 * @summary Patches container group profile with specified properties.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfilesPatch.json
 */
async function containerGroupProfilesPatch() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demoResource";
  const containerGroupProfileName = "demo1";
  const properties = {
    tags: { tag1key: "tag1Value", tag2key: "tag2Value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroupProfiles.patch(
    resourceGroupName,
    containerGroupProfileName,
    properties,
  );
  console.log(result);
}
