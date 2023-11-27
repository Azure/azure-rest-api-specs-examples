const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a managed private endpoint for a grafana resource.
 *
 * @summary Delete a managed private endpoint for a grafana resource.
 * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/ManagedPrivateEndpoints_Delete.json
 */
async function managedPrivateEndpointDelete() {
  const subscriptionId =
    process.env["DASHBOARD_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DASHBOARD_RESOURCE_GROUP"] || "myResourceGroup";
  const workspaceName = "myWorkspace";
  const managedPrivateEndpointName = "myMPEName";
  const credential = new DefaultAzureCredential();
  const client = new DashboardManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    managedPrivateEndpointName
  );
  console.log(result);
}
