const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all storages for a managedEnvironment.
 *
 * @summary Get all storages for a managedEnvironment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ManagedEnvironmentsStorages_List.json
 */
async function listEnvironmentsStoragesBySubscription() {
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = "examplerg";
  const environmentName = "managedEnv";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironmentsStorages.list(resourceGroupName, environmentName);
  console.log(result);
}

listEnvironmentsStoragesBySubscription().catch(console.error);
