const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the config reference app settings and status of an app
 *
 * @summary Description for Gets the config reference app settings and status of an app
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetKeyVaultReferencesForAppSettings.json
 */
async function getAzureKeyVaultReferencesForAppSettings() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "testc6282";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webApps.listAppSettingsKeyVaultReferences(
    resourceGroupName,
    name
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
