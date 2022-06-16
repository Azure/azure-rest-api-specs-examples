const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists ExpressRouteConnections.
 *
 * @summary Lists ExpressRouteConnections.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteConnectionList.json
 */
async function expressRouteConnectionList() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "expressRouteGatewayName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteConnections.list(
    resourceGroupName,
    expressRouteGatewayName
  );
  console.log(result);
}

expressRouteConnectionList().catch(console.error);
