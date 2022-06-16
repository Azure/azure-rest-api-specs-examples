const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get all deleted apps for a subscription at location
 *
 * @summary Description for Get all deleted apps for a subscription at location
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListDeletedWebAppsByLocation.json
 */
async function listDeletedWebAppByLocation() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const location = "West US 2";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedWebApps.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listDeletedWebAppByLocation().catch(console.error);
