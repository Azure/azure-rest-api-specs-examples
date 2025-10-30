const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a dashboard for Grafana resource.
 *
 * @summary delete a dashboard for Grafana resource.
 * x-ms-original-file: 2025-08-01/Dashboard_Delete.json
 */
async function dashboardDelete() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DashboardManagementClient(credential, subscriptionId);
  await client.managedDashboards.delete("myResourceGroup", "myDashboard");
}
