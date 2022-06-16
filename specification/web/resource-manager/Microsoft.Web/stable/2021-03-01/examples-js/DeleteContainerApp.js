const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Delete a Container App.
 *
 * @summary Description for Delete a Container App.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DeleteContainerApp.json
 */
async function deleteContainerApp() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testWorkerApp0";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.containerApps.beginDeleteAndWait(resourceGroupName, name);
  console.log(result);
}

deleteContainerApp().catch(console.error);
