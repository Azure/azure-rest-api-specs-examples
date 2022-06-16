const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Stop ongoing capturing network packets for the site.
 *
 * @summary Description for Stop ongoing capturing network packets for the site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StopWebSiteNetworkTrace.json
 */
async function stopACurrentlyRunningNetworkTraceOperationForASite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "SampleApp";
  const slot = "Production";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.stopWebSiteNetworkTraceSlot(resourceGroupName, name, slot);
  console.log(result);
}

stopACurrentlyRunningNetworkTraceOperationForASite().catch(console.error);
