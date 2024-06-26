const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the status of a private endpoint connection with the specified name
 *
 * @summary Update the status of a private endpoint connection with the specified name
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/PrivateEndpointConnectionsUpdate.json
 */
async function updateAPrivateEndpointConnection() {
  const subscriptionId = "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = "myResourceGroup";
  const workspaceName = "myWorkspace";
  const privateEndpointConnectionName = "myWorkspace.23456789-1111-1111-1111-111111111111";
  const privateEndpointConnection = {
    properties: {
      privateLinkServiceConnectionState: {
        description: "Approved by databricksadmin@contoso.com",
        status: "Approved",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDatabricksManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateAndWait(
    resourceGroupName,
    workspaceName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}

updateAPrivateEndpointConnection().catch(console.error);
