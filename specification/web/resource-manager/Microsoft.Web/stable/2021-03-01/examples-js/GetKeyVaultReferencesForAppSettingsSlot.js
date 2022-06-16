const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the config reference app settings and status of an app
 *
 * @summary Description for Gets the config reference app settings and status of an app
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetKeyVaultReferencesForAppSettingsSlot.json
 */
async function getAzureKeyVaultReferencesForAppSettingsForSlot() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "testc6282";
  const slot = "stage";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webApps.listAppSettingsKeyVaultReferencesSlot(
    resourceGroupName,
    name,
    slot
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAzureKeyVaultReferencesForAppSettingsForSlot().catch(console.error);
