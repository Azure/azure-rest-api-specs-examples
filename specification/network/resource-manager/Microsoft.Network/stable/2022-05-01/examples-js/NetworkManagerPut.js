const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Network Manager.
 *
 * @summary Creates or updates a Network Manager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerPut.json
 */
async function putNetworkManager() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "rg1";
  const networkManagerName = "TestNetworkManager";
  const parameters = {
    description: "My Test Network Manager",
    networkManagerScopeAccesses: ["Connectivity"],
    networkManagerScopes: {
      managementGroups: ["/Microsoft.Management/testmg"],
      subscriptions: ["/subscriptions/00000000-0000-0000-0000-000000000000"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkManagers.createOrUpdate(
    resourceGroupName,
    networkManagerName,
    parameters
  );
  console.log(result);
}

putNetworkManager().catch(console.error);
