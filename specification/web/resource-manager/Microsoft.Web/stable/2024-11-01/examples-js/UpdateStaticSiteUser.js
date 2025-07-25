const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Updates a user entry with the listed roles
 *
 * @summary Description for Updates a user entry with the listed roles
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateStaticSiteUser.json
 */
async function createOrUpdateAUserForAStaticSite() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "rg";
  const name = "testStaticSite0";
  const authprovider = "aad";
  const userid = "1234";
  const staticSiteUserEnvelope = {
    roles: "contributor",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.staticSites.updateStaticSiteUser(
    resourceGroupName,
    name,
    authprovider,
    userid,
    staticSiteUserEnvelope,
  );
  console.log(result);
}
