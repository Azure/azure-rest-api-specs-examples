const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified virtual network Gateway connection.
 *
 * @summary Deletes the specified virtual network Gateway connection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkGatewayConnectionDelete.json
 */
async function deleteVirtualNetworkGatewayConnection() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayConnectionName = "conn1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGatewayConnections.beginDeleteAndWait(
    resourceGroupName,
    virtualNetworkGatewayConnectionName
  );
  console.log(result);
}

deleteVirtualNetworkGatewayConnection().catch(console.error);
