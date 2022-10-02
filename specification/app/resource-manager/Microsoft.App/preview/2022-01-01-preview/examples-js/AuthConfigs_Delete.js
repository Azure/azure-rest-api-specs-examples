const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Delete a Container App AuthConfig.
 *
 * @summary Description for Delete a Container App AuthConfig.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/AuthConfigs_Delete.json
 */
async function deleteContainerAppAuthConfig() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const name = "current";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsAuthConfigs.delete(
    resourceGroupName,
    containerAppName,
    name
  );
  console.log(result);
}

deleteContainerAppAuthConfig().catch(console.error);
