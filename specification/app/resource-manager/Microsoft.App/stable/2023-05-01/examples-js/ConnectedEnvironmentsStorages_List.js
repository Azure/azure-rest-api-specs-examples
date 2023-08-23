const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all storages for a connectedEnvironment.
 *
 * @summary Get all storages for a connectedEnvironment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ConnectedEnvironmentsStorages_List.json
 */
async function listEnvironmentsStoragesBySubscription() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const connectedEnvironmentName = "managedEnv";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.connectedEnvironmentsStorages.list(
    resourceGroupName,
    connectedEnvironmentName
  );
  console.log(result);
}
