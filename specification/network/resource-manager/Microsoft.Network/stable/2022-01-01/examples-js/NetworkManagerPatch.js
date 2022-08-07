const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch NetworkManager.
 *
 * @summary Patch NetworkManager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerPatch.json
 */
async function networkManagesPatch() {
  const subscriptionId = "subscriptionC";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkManagers.patch(
    resourceGroupName,
    networkManagerName,
    parameters
  );
  console.log(result);
}

networkManagesPatch().catch(console.error);
