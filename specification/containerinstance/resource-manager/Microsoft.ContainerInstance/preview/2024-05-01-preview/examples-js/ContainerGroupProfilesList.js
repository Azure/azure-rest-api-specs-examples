const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of container group profiles in the specified subscription. This operation returns properties of each container group profile including containers, image registry credentials, restart policy, IP address type, OS type,volumes,current revision number, etc.
 *
 * @summary Get a list of container group profiles in the specified subscription. This operation returns properties of each container group profile including containers, image registry credentials, restart policy, IP address type, OS type,volumes,current revision number, etc.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfilesList.json
 */
async function containerGroupProfilesList() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerGroupProfiles.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
