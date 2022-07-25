const { Portal } = require("@azure/arm-portal");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Dashboard.
 *
 * @summary Gets the Dashboard.
 * x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/getDashboard.json
 */
async function getADashboard() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "testRG";
  const dashboardName = "testDashboard";
  const credential = new DefaultAzureCredential();
  const client = new Portal(credential, subscriptionId);
  const result = await client.dashboards.get(resourceGroupName, dashboardName);
  console.log(result);
}

getADashboard().catch(console.error);
