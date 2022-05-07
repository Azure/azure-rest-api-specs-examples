Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.1/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the quota and actual usage of the CDN profiles under the given subscription.
 *
 * @summary Check the quota and actual usage of the CDN profiles under the given subscription.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/ResourceUsage_List.json
 */
async function resourceUsageList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceUsageOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

resourceUsageList().catch(console.error);
```
