const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container groups with specified configurations.
 *
 * @summary Create or update container groups with specified configurations.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupCreateOrUpdateStandbyPool.json
 */
async function containerGroupCreateOrUpdateWithStandbyPool() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupName = "demo1";
  const containerGroup = {
    containerGroupProfile: {
      id: "/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/democgp",
      revision: 1,
    },
    containers: [{ name: "demo1", configMap: { keyValuePairs: { newkey: "value" } } }],
    location: "west us",
    standbyPoolProfile: {
      id: "/subscriptions/subid/resourceGroups/demo/providers/Microsoft.StandbyPool/standbyContainerGroupPools/demopool",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    containerGroupName,
    containerGroup,
  );
  console.log(result);
}
