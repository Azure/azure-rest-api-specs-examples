```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the configurationPolicyGroups in a resource group for a vpnServerConfiguration.
 *
 * @summary Lists all the configurationPolicyGroups in a resource group for a vpnServerConfiguration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ConfigurationPolicyGroupListByVpnServerConfiguration.json
 */
async function configurationPolicyGroupListByVpnServerConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnServerConfigurationName = "vpnServerConfiguration1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationPolicyGroups.listByVpnServerConfiguration(
    resourceGroupName,
    vpnServerConfigurationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

configurationPolicyGroupListByVpnServerConfiguration().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
