const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets the details of the user provided function apps registered with a static site build
 *
 * @summary Description for Gets the details of the user provided function apps registered with a static site build
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetUserProvidedFunctionAppsForStaticSiteBuild.json
 */
async function getDetailsOfTheUserProvidedFunctionAppsRegisteredWithAStaticSiteBuild() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const environmentName = "default";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.staticSites.listUserProvidedFunctionAppsForStaticSiteBuild(
    resourceGroupName,
    name,
    environmentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getDetailsOfTheUserProvidedFunctionAppsRegisteredWithAStaticSiteBuild().catch(console.error);
