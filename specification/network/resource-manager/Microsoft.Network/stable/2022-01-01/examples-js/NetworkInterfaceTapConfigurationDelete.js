const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified tap configuration from the NetworkInterface.
 *
 * @summary Deletes the specified tap configuration from the NetworkInterface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkInterfaceTapConfigurationDelete.json
 */
async function deleteTapConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkInterfaceName = "test-networkinterface";
  const tapConfigurationName = "test-tapconfiguration";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaceTapConfigurations.beginDeleteAndWait(
    resourceGroupName,
    networkInterfaceName,
    tapConfigurationName
  );
  console.log(result);
}

deleteTapConfiguration().catch(console.error);
