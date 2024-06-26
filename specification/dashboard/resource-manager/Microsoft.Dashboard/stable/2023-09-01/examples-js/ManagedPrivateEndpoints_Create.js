const { DashboardManagementClient } = require("@azure/arm-dashboard");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a managed private endpoint for a grafana resource.
 *
 * @summary Create or update a managed private endpoint for a grafana resource.
 * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2023-09-01/examples/ManagedPrivateEndpoints_Create.json
 */
async function managedPrivateEndpointCreate() {
  const subscriptionId =
    process.env["DASHBOARD_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DASHBOARD_RESOURCE_GROUP"] || "myResourceGroup";
  const workspaceName = "myWorkspace";
  const managedPrivateEndpointName = "myMPEName";
  const requestBodyParameters = {
    groupIds: ["grafana"],
    location: "West US",
    privateLinkResourceId:
      "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-000000000000/resourceGroups/xx-rg/providers/Microsoft.Kusto/Clusters/sampleKustoResource",
    privateLinkResourceRegion: "West US",
    privateLinkServiceUrl: "my-self-hosted-influxdb.westus.mydomain.com",
    requestMessage: "Example Request Message",
  };
  const credential = new DefaultAzureCredential();
  const client = new DashboardManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.beginCreateAndWait(
    resourceGroupName,
    workspaceName,
    managedPrivateEndpointName,
    requestBodyParameters
  );
  console.log(result);
}
