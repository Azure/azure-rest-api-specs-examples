const { AzureDatabricksManagementClient } = require("@azure/arm-databricks");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the status of a private endpoint connection with the specified name
 *
 * @summary Update the status of a private endpoint connection with the specified name
 * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/PrivateEndpointConnectionsUpdate.json
 */
async function updateAPrivateEndpointConnection() {
  const subscriptionId =
    process.env["DATABRICKS_SUBSCRIPTION_ID"] || "11111111-1111-1111-1111-111111111111";
  const resourceGroupName = process.env["DATABRICKS_RESOURCE_GROUP"] || "myResourceGroup";
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
