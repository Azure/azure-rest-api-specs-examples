const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Gets all scale-out instances of an app.
 *
 * @summary Description for Gets all scale-out instances of an app.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/GetSiteInstanceInfo.json
 */
async function getSiteInstanceInfo() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "tests346";
  const instanceId = "134987120";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.getInstanceInfo(resourceGroupName, name, instanceId);
  console.log(result);
}

getSiteInstanceInfo().catch(console.error);
