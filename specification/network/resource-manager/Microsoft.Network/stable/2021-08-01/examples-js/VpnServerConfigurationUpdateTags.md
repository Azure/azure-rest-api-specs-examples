Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates VpnServerConfiguration tags.
 *
 * @summary Updates VpnServerConfiguration tags.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnServerConfigurationUpdateTags.json
 */
async function vpnServerConfigurationUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const vpnServerConfigurationName = "vpnServerConfiguration1";
  const vpnServerConfigurationParameters = {
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnServerConfigurations.updateTags(
    resourceGroupName,
    vpnServerConfigurationName,
    vpnServerConfigurationParameters
  );
  console.log(result);
}

vpnServerConfigurationUpdate().catch(console.error);
```
