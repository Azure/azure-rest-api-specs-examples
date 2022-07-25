const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a private endpoint connection.
 *
 * @summary Delete a private endpoint connection.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeletePrivateEndpointConnection.json
 */
async function deletePrivateEndpointConnection() {
  const subscriptionId = "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const privateEndpointConnectionName = "ExamplePrivateEndpointConnection";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    privateEndpointConnectionName
  );
  console.log(result);
}

deletePrivateEndpointConnection().catch(console.error);
