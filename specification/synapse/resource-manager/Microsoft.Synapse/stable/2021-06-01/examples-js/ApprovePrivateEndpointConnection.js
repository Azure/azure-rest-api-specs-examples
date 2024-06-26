const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Approve or reject a private endpoint connection.
 *
 * @summary Approve or reject a private endpoint connection.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ApprovePrivateEndpointConnection.json
 */
async function approvePrivateEndpointConnection() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const workspaceName = "ExampleWorkspace";
  const privateEndpointConnectionName = "ExamplePrivateEndpointConnection";
  const request = {
    privateLinkServiceConnectionState: {
      description: "Approved by abc@example.com",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateAndWait(
    resourceGroupName,
    workspaceName,
    privateEndpointConnectionName,
    request
  );
  console.log(result);
}
