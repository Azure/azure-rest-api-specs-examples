const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network group.
 *
 * @summary Creates or updates a network group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerGroupPut.json
 */
async function networkGroupsPut() {
  const subscriptionId = "subscriptionC";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const networkGroupName = "testNetworkGroup";
  const parameters = { description: "A sample group" };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkGroups.createOrUpdate(
    resourceGroupName,
    networkManagerName,
    networkGroupName,
    parameters
  );
  console.log(result);
}

networkGroupsPut().catch(console.error);
