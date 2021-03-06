const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of the specified private endpoint connection associated with the workspace.
 *
 * @summary Update the state of the specified private endpoint connection associated with the workspace.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/privatelink/WorkspaceCreatePrivateEndpointConnection.json
 */
async function workspacePrivateEndpointConnectionCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const privateEndpointConnectionName = "myConnection";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.workspacePrivateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}

workspacePrivateEndpointConnectionCreateOrUpdate().catch(console.error);
