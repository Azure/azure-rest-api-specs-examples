const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of the specified private endpoint connection associated with the workspace.
 *
 * @summary Update the state of the specified private endpoint connection associated with the workspace.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/privatelink/WorkspaceCreatePrivateEndpointConnection.json
 */
async function workspacePrivateEndpointConnectionCreateOrUpdate() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "testRG";
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
    properties,
  );
  console.log(result);
}
