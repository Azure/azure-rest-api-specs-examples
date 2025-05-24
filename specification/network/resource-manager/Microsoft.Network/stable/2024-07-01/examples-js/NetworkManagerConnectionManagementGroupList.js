const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all network manager connections created by this management group.
 *
 * @summary List all network manager connections created by this management group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerConnectionManagementGroupList.json
 */
async function listManagementGroupNetworkManagerConnection() {
  const managementGroupId = "managementGroupA";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.managementGroupNetworkManagerConnections.list(
    managementGroupId,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
