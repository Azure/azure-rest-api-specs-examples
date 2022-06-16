const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Container App AuthConfigs in a given resource group.
 *
 * @summary Get the Container App AuthConfigs in a given resource group.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/AuthConfigs_ListByContainer.json
 */
async function listAuthConfigsByContainerApps() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerAppsAuthConfigs.listByContainerApp(
    resourceGroupName,
    containerAppName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAuthConfigsByContainerApps().catch(console.error);
