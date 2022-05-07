Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-frontdoor_5.0.1/sdk/frontdoor/arm-frontdoor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the protection policies within a resource group.
 *
 * @summary Lists all of the protection policies within a resource group.
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-11-01/examples/WafListPolicies.json
 */
async function listPoliciesInAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policies.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listPoliciesInAResourceGroup().catch(console.error);
```
