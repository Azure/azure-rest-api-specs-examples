const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the specified container group profile in the specified subscription and resource group. The operation does not delete other resources provided by the user, such as volumes.
 *
 * @summary Delete the specified container group profile in the specified subscription and resource group. The operation does not delete other resources provided by the user, such as volumes.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfilesDelete.json
 */
async function containerGroupProfilesDelete() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupProfileName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroupProfiles.delete(
    resourceGroupName,
    containerGroupProfileName,
  );
  console.log(result);
}
