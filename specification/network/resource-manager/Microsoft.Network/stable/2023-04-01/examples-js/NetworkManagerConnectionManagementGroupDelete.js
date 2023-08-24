const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete specified pending connection created by this management group.
 *
 * @summary Delete specified pending connection created by this management group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/NetworkManagerConnectionManagementGroupDelete.json
 */
async function deleteManagementGroupNetworkManagerConnection() {
  const managementGroupId = "managementGroupA";
  const networkManagerConnectionName = "TestNMConnection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential);
  const result = await client.managementGroupNetworkManagerConnections.delete(
    managementGroupId,
    networkManagerConnectionName
  );
  console.log(result);
}
