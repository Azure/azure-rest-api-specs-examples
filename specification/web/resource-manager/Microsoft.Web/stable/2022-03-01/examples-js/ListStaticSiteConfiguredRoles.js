const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Lists the roles configured for the static site.
 *
 * @summary Description for Lists the roles configured for the static site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/ListStaticSiteConfiguredRoles.json
 */
async function listsTheConfiguredRolesForAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.listStaticSiteConfiguredRoles(resourceGroupName, name);
  console.log(result);
}

listsTheConfiguredRolesForAStaticSite().catch(console.error);
