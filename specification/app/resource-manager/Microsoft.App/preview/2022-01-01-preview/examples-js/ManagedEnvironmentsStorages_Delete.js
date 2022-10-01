const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete storage for a managedEnvironment.
 *
 * @summary Delete storage for a managedEnvironment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ManagedEnvironmentsStorages_Delete.json
 */
async function listEnvironmentsStoragesBySubscription() {
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = "examplerg";
  const envName = "managedEnv";
  const name = "jlaw-demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironmentsStorages.delete(resourceGroupName, envName, name);
  console.log(result);
}

listEnvironmentsStoragesBySubscription().catch(console.error);
