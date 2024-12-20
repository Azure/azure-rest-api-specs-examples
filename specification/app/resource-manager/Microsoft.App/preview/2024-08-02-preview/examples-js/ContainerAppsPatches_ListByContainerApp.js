const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Container Apps Patch resources by ContainerApp.
 *
 * @summary List Container Apps Patch resources by ContainerApp.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerAppsPatches_ListByContainerApp.json
 */
async function containerAppsPatchesListByContainerApp0() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "test-app";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerAppsPatches.listByContainerApp(
    resourceGroupName,
    containerAppName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
