Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the specified default network security rule.
 *
 * @summary Get the specified default network security rule.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/DefaultSecurityRuleGet.json
 */
async function defaultSecurityRuleGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const networkSecurityGroupName = "nsg1";
  const defaultSecurityRuleName = "AllowVnetInBound";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.defaultSecurityRules.get(
    resourceGroupName,
    networkSecurityGroupName,
    defaultSecurityRuleName
  );
  console.log(result);
}

defaultSecurityRuleGet().catch(console.error);
```
