const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a network manager.
 *
 * @summary Deletes a network manager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerDelete.json
 */
async function networkManagersDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const force = false;
  const options = { force };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkManagers.beginDeleteAndWait(
    resourceGroupName,
    networkManagerName,
    options
  );
  console.log(result);
}

networkManagersDelete().catch(console.error);
