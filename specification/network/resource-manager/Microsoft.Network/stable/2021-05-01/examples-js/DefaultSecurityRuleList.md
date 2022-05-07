Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all default security rules in a network security group.
 *
 * @summary Gets all default security rules in a network security group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/DefaultSecurityRuleList.json
 */
async function defaultSecurityRuleList() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const networkSecurityGroupName = "nsg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.defaultSecurityRules.list(
    resourceGroupName,
    networkSecurityGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

defaultSecurityRuleList().catch(console.error);
```
