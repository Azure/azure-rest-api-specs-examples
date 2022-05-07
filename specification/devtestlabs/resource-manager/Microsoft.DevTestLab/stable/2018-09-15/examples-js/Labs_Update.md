Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Allows modifying tags of labs. All other properties will be ignored.
 *
 * @summary Allows modifying tags of labs. All other properties will be ignored.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_Update.json
 */
async function labsUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const lab = {};
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.labs.update(resourceGroupName, name, lab);
  console.log(result);
}

labsUpdate().catch(console.error);
```
