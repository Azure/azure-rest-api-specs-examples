const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a AuthConfig of a Container App.
 *
 * @summary Get a AuthConfig of a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/AuthConfigs_Get.json
 */
async function getContainerAppAuthConfig() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const authConfigName = "current";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsAuthConfigs.get(
    resourceGroupName,
    containerAppName,
    authConfigName
  );
  console.log(result);
}

getContainerAppAuthConfig().catch(console.error);
