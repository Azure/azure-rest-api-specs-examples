const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the specified container group in the specified subscription and resource group. The operation does not delete other resources provided by the user, such as volumes.
 *
 * @summary Delete the specified container group in the specified subscription and resource group. The operation does not delete other resources provided by the user, such as volumes.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsDelete.json
 */
async function containerGroupsDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.beginDeleteAndWait(
    resourceGroupName,
    containerGroupName
  );
  console.log(result);
}

containerGroupsDelete().catch(console.error);
