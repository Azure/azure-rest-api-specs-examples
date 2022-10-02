const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get storage for a managedEnvironment.
 *
 * @summary Get storage for a managedEnvironment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ManagedEnvironmentsStorages_Get.json
 */
async function getAEnvironmentsStoragePropertiesBySubscription() {
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = "examplerg";
  const envName = "managedEnv";
  const name = "jlaw-demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironmentsStorages.get(resourceGroupName, envName, name);
  console.log(result);
}

getAEnvironmentsStoragePropertiesBySubscription().catch(console.error);
