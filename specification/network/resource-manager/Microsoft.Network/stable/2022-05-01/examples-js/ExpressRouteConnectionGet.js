const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified ExpressRouteConnection.
 *
 * @summary Gets the specified ExpressRouteConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteConnectionGet.json
 */
async function expressRouteConnectionGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "expressRouteGatewayName";
  const connectionName = "connectionName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteConnections.get(
    resourceGroupName,
    expressRouteGatewayName,
    connectionName
  );
  console.log(result);
}

expressRouteConnectionGet().catch(console.error);
