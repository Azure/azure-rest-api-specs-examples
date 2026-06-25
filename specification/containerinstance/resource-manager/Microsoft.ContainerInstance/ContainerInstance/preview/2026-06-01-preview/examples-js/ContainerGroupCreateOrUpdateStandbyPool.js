const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update container groups with specified configurations.
 *
 * @summary create or update container groups with specified configurations.
 * x-ms-original-file: 2026-06-01-preview/ContainerGroupCreateOrUpdateStandbyPool.json
 */
async function containerGroupCreateOrUpdateWithStandbyPool() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.createOrUpdate("demo", "demo1", {
    location: "west us",
    containerGroupProfile: {
      id: "/subscriptions/subid/resourceGroups/demo/providers/Microsoft.ContainerInstance/containerGroupProfiles/democgp",
      revision: 1,
    },
    containers: [{ name: "demo1", configMap: { keyValuePairs: { Newkey: "value" } } }],
    standbyPoolProfile: {
      id: "/subscriptions/subid/resourceGroups/demo/providers/Microsoft.StandbyPool/standbyContainerGroupPools/demopool",
    },
  });
  console.log(result);
}
