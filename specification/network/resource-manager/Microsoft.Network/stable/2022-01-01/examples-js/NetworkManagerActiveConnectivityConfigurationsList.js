const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists active connectivity configurations in a network manager.
 *
 * @summary Lists active connectivity configurations in a network manager.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerActiveConnectivityConfigurationsList.json
 */
async function listActiveConnectivityConfigurations() {
  const subscriptionId = "subscriptionA";
  const resourceGroupName = "myResourceGroup";
  const networkManagerName = "testNetworkManager";
  const parameters = {
    regions: ["westus"],
    skipToken: "fakeSkipTokenCode",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.listActiveConnectivityConfigurations(
    resourceGroupName,
    networkManagerName,
    parameters
  );
  console.log(result);
}

listActiveConnectivityConfigurations().catch(console.error);
