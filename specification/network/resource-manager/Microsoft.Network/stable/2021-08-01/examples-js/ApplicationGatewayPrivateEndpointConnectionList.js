const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all private endpoint connections on an application gateway.
 *
 * @summary Lists all private endpoint connections on an application gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ApplicationGatewayPrivateEndpointConnectionList.json
 */
async function listsAllPrivateEndpointConnectionsOnApplicationGateway() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const applicationGatewayName = "appgw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationGatewayPrivateEndpointConnections.list(
    resourceGroupName,
    applicationGatewayName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllPrivateEndpointConnectionsOnApplicationGateway().catch(console.error);
