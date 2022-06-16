const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private endpoint connection on application gateway.
 *
 * @summary Deletes the specified private endpoint connection on application gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayPrivateEndpointConnectionDelete.json
 */
async function deleteApplicationGatewayPrivateEndpointConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const applicationGatewayName = "appgw";
  const connectionName = "connection1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGatewayPrivateEndpointConnections.beginDeleteAndWait(
    resourceGroupName,
    applicationGatewayName,
    connectionName
  );
  console.log(result);
}

deleteApplicationGatewayPrivateEndpointConnection().catch(console.error);
