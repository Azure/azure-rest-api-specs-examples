const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a network manager connectivity configuration, specified by the resource group, network manager name, and connectivity configuration name
 *
 * @summary Deletes a network manager connectivity configuration, specified by the resource group, network manager name, and connectivity configuration name
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerConnectivityConfigurationDelete.json
 */
async function connectivityConfigurationsDelete() {
  const subscriptionId = "subscriptionA";
  const resourceGroupName = "myResourceGroup";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestConnectivityConfig";
  const force = false;
  const options = { force };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.connectivityConfigurations.beginDeleteAndWait(
    resourceGroupName,
    networkManagerName,
    configurationName,
    options
  );
  console.log(result);
}

connectivityConfigurationsDelete().catch(console.error);
