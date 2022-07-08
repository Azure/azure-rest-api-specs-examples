const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the details of the user provided function app registered with a static site
 *
 * @summary Description for Gets the details of the user provided function app registered with a static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetUserProvidedFunctionAppForStaticSite.json
 */
async function getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const functionAppName = "testFunctionApp";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.getUserProvidedFunctionAppForStaticSite(
    resourceGroupName,
    name,
    functionAppName
  );
  console.log(result);
}

getDetailsOfTheUserProvidedFunctionAppRegisteredWithAStaticSite().catch(console.error);
