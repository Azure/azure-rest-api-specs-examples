const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all private link resources information for this grafana resource
 *
 * @summary List all private link resources information for this grafana resource
 * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/PrivateLinkResources_List.json
 */
async function privateLinkResourcesList() {
  const subscriptionId =
    process.env["DASHBOARD_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DASHBOARD_RESOURCE_GROUP"] || "myResourceGroup";
  const workspaceName = "myWorkspace";
  const credential = new DefaultAzureCredential();
  const client = new DashboardManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkResources.list(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
