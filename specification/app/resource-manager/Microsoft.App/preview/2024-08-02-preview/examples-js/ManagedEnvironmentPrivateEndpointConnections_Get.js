const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a private endpoint connection for a given managed environment.
 *
 * @summary Get a private endpoint connection for a given managed environment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ManagedEnvironmentPrivateEndpointConnections_Get.json
 */
async function getAPrivateEndpointConnectionByManagedEnvironment() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "managedEnv";
  const privateEndpointConnectionName = "jlaw-demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironmentPrivateEndpointConnections.get(
    resourceGroupName,
    environmentName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
