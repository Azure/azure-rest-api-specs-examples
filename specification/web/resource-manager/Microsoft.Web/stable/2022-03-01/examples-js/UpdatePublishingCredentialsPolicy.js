const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Updates whether user publishing credentials are allowed on the site or not.
 *
 * @summary Description for Updates whether user publishing credentials are allowed on the site or not.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/UpdatePublishingCredentialsPolicy.json
 */
async function updateScmAllowed() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testSite";
  const csmPublishingAccessPoliciesEntity = {
    allow: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.updateScmAllowed(
    resourceGroupName,
    name,
    csmPublishingAccessPoliciesEntity
  );
  console.log(result);
}

updateScmAllowed().catch(console.error);
