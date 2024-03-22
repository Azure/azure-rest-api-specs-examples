const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get storage for a managedEnvironment.
 *
 * @summary Get storage for a managedEnvironment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ManagedEnvironmentsStorages_Get_NfsAzureFile.json
 */
async function getAEnvironmentsStorageForNfsAzureFile() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "managedEnv";
  const storageName = "jlaw-demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironmentsStorages.get(
    resourceGroupName,
    environmentName,
    storageName,
  );
  console.log(result);
}
