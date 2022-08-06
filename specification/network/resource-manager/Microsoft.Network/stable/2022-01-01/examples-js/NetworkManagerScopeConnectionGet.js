const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get specified scope connection created by this Network Manager.
 *
 * @summary Get specified scope connection created by this Network Manager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerScopeConnectionGet.json
 */
async function getNetworkManagerScopeConnection() {
  const subscriptionId = "subscriptionC";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const scopeConnectionName = "TestScopeConnection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.scopeConnections.get(
    resourceGroupName,
    networkManagerName,
    scopeConnectionName
  );
  console.log(result);
}

getNetworkManagerScopeConnection().catch(console.error);
