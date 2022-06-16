const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified private endpoint connection on application gateway.
 *
 * @summary Updates the specified private endpoint connection on application gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayPrivateEndpointConnectionUpdate.json
 */
async function updateApplicationGatewayPrivateEndpointConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const applicationGatewayName = "appgw";
  const connectionName = "connection1";
  const parameters = {
    name: "connection1",
    privateEndpoint: {
      id: "/subscriptions/subId2/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe",
    },
    privateLinkServiceConnectionState: {
      description: "approved it for some reason.",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGatewayPrivateEndpointConnections.beginUpdateAndWait(
    resourceGroupName,
    applicationGatewayName,
    connectionName,
    parameters
  );
  console.log(result);
}

updateApplicationGatewayPrivateEndpointConnection().catch(console.error);
