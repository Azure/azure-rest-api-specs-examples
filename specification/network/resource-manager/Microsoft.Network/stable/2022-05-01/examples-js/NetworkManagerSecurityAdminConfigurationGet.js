const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves a network manager security admin configuration.
 *
 * @summary Retrieves a network manager security admin configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerSecurityAdminConfigurationGet.json
 */
async function getSecurityAdminConfigurations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestSecurityConfig";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityAdminConfigurations.get(
    resourceGroupName,
    networkManagerName,
    configurationName
  );
  console.log(result);
}

getSecurityAdminConfigurations().catch(console.error);
