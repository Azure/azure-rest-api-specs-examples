const createNetworkManagementClient = require("@azure-rest/arm-network").default;
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates scope connection from Network Manager
 *
 * @summary Creates or updates scope connection from Network Manager
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerScopeConnectionPut.json
 */
async function createOrUpdateNetworkManagerScopeConnection() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const scopeConnectionName = "TestScopeConnection";
  const options = {
    body: {
      properties: {
        description: "This is a scope connection to a cross tenant subscription.",
        resourceId: "subscriptions/f0dc2b34-dfad-40e4-83e0-2309fed8d00b",
        tenantId: "6babcaad-604b-40ac-a9d7-9fd97c0b779f",
      },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}",
      subscriptionId,
      resourceGroupName,
      networkManagerName,
      scopeConnectionName,
    )
    .put(options);
  console.log(result);
}

createOrUpdateNetworkManagerScopeConnection().catch(console.error);
