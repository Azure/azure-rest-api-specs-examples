Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-batch_7.1.0/sdk/batch/arm-batch/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Batch supported Cloud Service VM sizes available at the given location.
 *
 * @summary Gets the list of Batch supported Cloud Service VM sizes available at the given location.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/LocationListCloudServiceSkus.json
 */
async function locationListCloudServiceSkus() {
  const subscriptionId = "subid";
  const locationName = "japaneast";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.location.listSupportedCloudServiceSkus(locationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

locationListCloudServiceSkus().catch(console.error);
```
