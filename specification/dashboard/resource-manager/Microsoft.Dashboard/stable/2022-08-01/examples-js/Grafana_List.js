const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all resources of workspaces for Grafana under the specified subscription.
 *
 * @summary List all resources of workspaces for Grafana under the specified subscription.
 * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/Grafana_List.json
 */
async function grafanaList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DashboardManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.grafana.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

grafanaList().catch(console.error);
