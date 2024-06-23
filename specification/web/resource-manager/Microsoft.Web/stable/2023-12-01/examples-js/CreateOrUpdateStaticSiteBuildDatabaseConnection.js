const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Create or update a database connection for a static site build
 *
 * @summary Description for Create or update a database connection for a static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/CreateOrUpdateStaticSiteBuildDatabaseConnection.json
 */
async function createOrUpdateADatabaseConnectionForAStaticSiteBuild() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testStaticSite0";
  const environmentName = "default";
  const databaseConnectionName = "default";
  const databaseConnectionRequestEnvelope = {
    connectionIdentity: "SystemAssigned",
    connectionString:
      "AccountEndpoint=https://exampleDatabaseName.documents.azure.com:443/;Database=mydb;",
    region: "West US 2",
    resourceId:
      "/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/databaseRG/providers/Microsoft.DocumentDB/databaseAccounts/exampleDatabaseName",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.createOrUpdateBuildDatabaseConnection(
    resourceGroupName,
    name,
    environmentName,
    databaseConnectionName,
    databaseConnectionRequestEnvelope,
  );
  console.log(result);
}
