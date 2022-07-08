const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets all static site builds for a particular static site.
 *
 * @summary Description for Gets all static site builds for a particular static site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetStaticSiteBuilds.json
 */
async function getAllBuildsForAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.staticSites.listStaticSiteBuilds(resourceGroupName, name)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllBuildsForAStaticSite().catch(console.error);
