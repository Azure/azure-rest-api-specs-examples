Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Tap configuration in the specified NetworkInterface.
 *
 * @summary Creates or updates a Tap configuration in the specified NetworkInterface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkInterfaceTapConfigurationCreate.json
 */
async function createNetworkInterfaceTapConfigurations() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const networkInterfaceName = "mynic";
  const tapConfigurationName = "tapconfiguration1";
  const tapConfigurationParameters = {
    virtualNetworkTap: {
      id: "/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkTaps/testvtap",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaceTapConfigurations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkInterfaceName,
    tapConfigurationName,
    tapConfigurationParameters
  );
  console.log(result);
}

createNetworkInterfaceTapConfigurations().catch(console.error);
```
