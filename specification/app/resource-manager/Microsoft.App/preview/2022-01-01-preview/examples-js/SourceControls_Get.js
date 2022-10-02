const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a SourceControl of a Container App.
 *
 * @summary Get a SourceControl of a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/SourceControls_Get.json
 */
async function getContainerAppSourceControl() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const name = "current";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsSourceControls.get(
    resourceGroupName,
    containerAppName,
    name
  );
  console.log(result);
}

getContainerAppSourceControl().catch(console.error);
