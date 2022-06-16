const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Returns whether Scm basic auth is allowed on the site or not.
 *
 * @summary Description for Returns whether Scm basic auth is allowed on the site or not.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetPublishingCredentialsPolicy.json
 */
async function getScmAllowed() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testSite";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.getScmAllowed(resourceGroupName, name);
  console.log(result);
}

getScmAllowed().catch(console.error);
