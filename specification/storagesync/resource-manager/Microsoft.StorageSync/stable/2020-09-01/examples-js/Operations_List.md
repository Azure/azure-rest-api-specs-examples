Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagesync_9.0.1/sdk/storagesync/arm-storagesync/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available Storage Sync Rest API operations.
 *
 * @summary Lists all of the available Storage Sync Rest API operations.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/Operations_List.json
 */
async function operationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.operations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

operationsList().catch(console.error);
```
