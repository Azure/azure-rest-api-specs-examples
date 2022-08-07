const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Network Connectivity Configuration, specified by the resource group, network manager name, and connectivity Configuration name
 *
 * @summary Gets a Network Connectivity Configuration, specified by the resource group, network manager name, and connectivity Configuration name
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectivityConfigurationGet.json
 */
async function connectivityConfigurationsGet() {
  const subscriptionId = "subscriptionA";
  const resourceGroupName = "myResourceGroup";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestConnectivityConfig";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectivityConfigurations.get(
    resourceGroupName,
    networkManagerName,
    configurationName
  );
  console.log(result);
}

connectivityConfigurationsGet().catch(console.error);
