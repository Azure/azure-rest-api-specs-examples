Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get disk.
 *
 * @summary Get disk.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_Get.json
 */
async function disksGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const userName = "@me";
  const name = "{diskName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.disks.get(resourceGroupName, labName, userName, name);
  console.log(result);
}

disksGet().catch(console.error);
```
