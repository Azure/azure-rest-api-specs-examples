const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Deletes the user entry from the static site.
 *
 * @summary Description for Deletes the user entry from the static site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/DeleteStaticSiteUser.json
 */
async function deleteAUserForAStaticSite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testStaticSite0";
  const authprovider = "aad";
  const userid = "1234";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.deleteStaticSiteUser(
    resourceGroupName,
    name,
    authprovider,
    userid
  );
  console.log(result);
}

deleteAUserForAStaticSite().catch(console.error);
