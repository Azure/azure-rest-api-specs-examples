const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all resources of workspaces for Grafana under the specified resource group.
 *
 * @summary List all resources of workspaces for Grafana under the specified resource group.
 * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/preview/2021-09-01-preview/examples/Grafana_ListByResourceGroup.json
 */
async function grafanaListByResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new DashboardManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.grafana.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

grafanaListByResourceGroup().catch(console.error);
