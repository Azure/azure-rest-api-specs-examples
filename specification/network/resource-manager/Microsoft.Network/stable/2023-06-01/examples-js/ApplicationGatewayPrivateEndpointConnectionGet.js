const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified private endpoint connection on application gateway.
 *
 * @summary Gets the specified private endpoint connection on application gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-06-01/examples/ApplicationGatewayPrivateEndpointConnectionGet.json
 */
async function getApplicationGatewayPrivateEndpointConnection() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const applicationGatewayName = "appgw";
  const connectionName = "connection1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGatewayPrivateEndpointConnections.get(
    resourceGroupName,
    applicationGatewayName,
    connectionName,
  );
  console.log(result);
}
