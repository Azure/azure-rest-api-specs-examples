const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Container App AuthConfig.
 *
 * @summary Delete a Container App AuthConfig.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/AuthConfigs_Delete.json
 */
async function deleteContainerAppAuthConfig() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const authConfigName = "current";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsAuthConfigs.delete(
    resourceGroupName,
    containerAppName,
    authConfigName
  );
  console.log(result);
}

deleteContainerAppAuthConfig().catch(console.error);
