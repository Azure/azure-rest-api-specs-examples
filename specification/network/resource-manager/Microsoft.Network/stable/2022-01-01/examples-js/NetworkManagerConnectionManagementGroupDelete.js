const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete specified pending connection created by this management group.
 *
 * @summary Delete specified pending connection created by this management group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectionManagementGroupDelete.json
 */
async function deleteManagementGroupNetworkManagerConnection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "managementGroupA";
  const networkManagerConnectionName = "TestNMConnection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.managementGroupNetworkManagerConnections.delete(
    managementGroupId,
    networkManagerConnectionName
  );
  console.log(result);
}

deleteManagementGroupNetworkManagerConnection().catch(console.error);
