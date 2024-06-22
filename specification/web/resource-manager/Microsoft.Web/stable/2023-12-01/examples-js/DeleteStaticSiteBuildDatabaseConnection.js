const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a database connection for a static site build
 *
 * @summary Delete a database connection for a static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/DeleteStaticSiteBuildDatabaseConnection.json
 */
async function deleteADatabaseConnectionFromAStaticSiteBuild() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testStaticSite0";
  const environmentName = "default";
  const databaseConnectionName = "default";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.deleteBuildDatabaseConnection(
    resourceGroupName,
    name,
    environmentName,
    databaseConnectionName,
  );
  console.log(result);
}
