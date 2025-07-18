const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Returns details of database connections for a static site build
 *
 * @summary Returns details of database connections for a static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetStaticSiteBuildDatabaseConnectionsWithDetails.json
 */
async function listFullDetailsOfDatabaseConnectionsForTheStaticSiteBuild() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testStaticSite0";
  const environmentName = "default";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.staticSites.listBuildDatabaseConnectionsWithDetails(
    resourceGroupName,
    name,
    environmentName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
