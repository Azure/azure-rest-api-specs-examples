const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a network manager security admin configuration.
 *
 * @summary Deletes a network manager security admin configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkManagerSecurityAdminConfigurationDelete.json
 */
async function deleteNetworkManagerSecurityAdminConfiguration() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestSecurityConfig";
  const force = false;
  const options = { force };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityAdminConfigurations.beginDeleteAndWait(
    resourceGroupName,
    networkManagerName,
    configurationName,
    options
  );
  console.log(result);
}

deleteNetworkManagerSecurityAdminConfiguration().catch(console.error);
