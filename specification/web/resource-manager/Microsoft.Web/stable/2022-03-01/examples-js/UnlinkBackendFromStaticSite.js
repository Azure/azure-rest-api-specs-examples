const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Unlink a backend from a static site
 *
 * @summary Unlink a backend from a static site
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/UnlinkBackendFromStaticSite.json
 */
async function unlinkABackendFromAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const linkedBackendName = "testBackend";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.unlinkBackend(resourceGroupName, name, linkedBackendName);
  console.log(result);
}

unlinkABackendFromAStaticSite().catch(console.error);
