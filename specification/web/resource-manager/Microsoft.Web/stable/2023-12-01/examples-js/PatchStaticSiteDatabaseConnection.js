const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Create or update a database connection for a static site
 *
 * @summary Description for Create or update a database connection for a static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/PatchStaticSiteDatabaseConnection.json
 */
async function patchADatabaseConnectionForAStaticSite() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testStaticSite0";
  const databaseConnectionName = "default";
  const databaseConnectionRequestEnvelope = {};
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.updateDatabaseConnection(
    resourceGroupName,
    name,
    databaseConnectionName,
    databaseConnectionRequestEnvelope,
  );
  console.log(result);
}
