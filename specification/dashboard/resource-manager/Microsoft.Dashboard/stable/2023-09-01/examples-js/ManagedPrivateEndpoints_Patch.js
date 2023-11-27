const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a managed private endpoint for an existing grafana resource.
 *
 * @summary Update a managed private endpoint for an existing grafana resource.
 * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/ManagedPrivateEndpoints_Patch.json
 */
async function managedPrivateEndpointsPatch() {
  const subscriptionId =
    process.env["DASHBOARD_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DASHBOARD_RESOURCE_GROUP"] || "myResourceGroup";
  const workspaceName = "myWorkspace";
  const managedPrivateEndpointName = "myMPEName";
  const requestBodyParameters = {
    tags: { environment: "Dev 2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new DashboardManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.beginUpdateAndWait(
    resourceGroupName,
    workspaceName,
    managedPrivateEndpointName,
    requestBodyParameters
  );
  console.log(result);
}
