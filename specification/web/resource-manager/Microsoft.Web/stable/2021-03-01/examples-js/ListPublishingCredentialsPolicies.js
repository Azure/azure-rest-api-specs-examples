const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Returns whether Scm basic auth is allowed and whether Ftp is allowed for a given site.
 *
 * @summary Description for Returns whether Scm basic auth is allowed and whether Ftp is allowed for a given site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListPublishingCredentialsPolicies.json
 */
async function listPublishingCredentialsPolicies() {
  const subscriptionId = "3fb8d758-2e2c-42e9-a528-a8acdfe87237";
  const resourceGroupName = "testrg123";
  const name = "testsite";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.webApps.listBasicPublishingCredentialsPolicies(
    resourceGroupName,
    name
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPublishingCredentialsPolicies().catch(console.error);
