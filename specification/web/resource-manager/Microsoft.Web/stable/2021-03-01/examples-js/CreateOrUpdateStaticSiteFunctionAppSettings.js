const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates or updates the function app settings of a static site.
 *
 * @summary Description for Creates or updates the function app settings of a static site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateStaticSiteFunctionAppSettings.json
 */
async function createsOrUpdatesTheFunctionAppSettingsOfAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const appSettings = {
    properties: { setting1: "someval", setting2: "someval2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.createOrUpdateStaticSiteFunctionAppSettings(
    resourceGroupName,
    name,
    appSettings
  );
  console.log(result);
}

createsOrUpdatesTheFunctionAppSettingsOfAStaticSite().catch(console.error);
