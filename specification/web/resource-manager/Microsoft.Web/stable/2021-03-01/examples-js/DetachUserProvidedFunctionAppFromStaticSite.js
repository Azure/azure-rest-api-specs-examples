const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Detach the user provided function app from the static site
 *
 * @summary Description for Detach the user provided function app from the static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/DetachUserProvidedFunctionAppFromStaticSite.json
 */
async function detachTheUserProvidedFunctionAppFromTheStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const functionAppName = "testFunctionApp";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.detachUserProvidedFunctionAppFromStaticSite(
    resourceGroupName,
    name,
    functionAppName
  );
  console.log(result);
}

detachTheUserProvidedFunctionAppFromTheStaticSite().catch(console.error);
