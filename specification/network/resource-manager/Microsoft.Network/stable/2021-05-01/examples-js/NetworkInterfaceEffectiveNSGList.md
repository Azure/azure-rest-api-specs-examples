Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all network security groups applied to a network interface.
 *
 * @summary Gets all network security groups applied to a network interface.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkInterfaceEffectiveNSGList.json
 */
async function listNetworkInterfaceEffectiveNetworkSecurityGroups() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkInterfaceName = "nic1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkInterfaces.beginListEffectiveNetworkSecurityGroupsAndWait(
    resourceGroupName,
    networkInterfaceName
  );
  console.log(result);
}

listNetworkInterfaceEffectiveNetworkSecurityGroups().catch(console.error);
```
