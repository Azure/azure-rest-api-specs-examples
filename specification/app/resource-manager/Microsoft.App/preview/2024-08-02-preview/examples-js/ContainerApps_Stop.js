const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Stop a container app
 *
 * @summary Stop a container app
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/ContainerApps_Stop.json
 */
async function stopContainerApp() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testWorkerApp0";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.beginStopAndWait(resourceGroupName, containerAppName);
  console.log(result);
}
