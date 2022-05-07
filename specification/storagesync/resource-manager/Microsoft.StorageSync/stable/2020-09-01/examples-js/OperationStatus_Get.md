Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storagesync_9.0.1/sdk/storagesync/arm-storagesync/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Operation status
 *
 * @summary Get Operation status
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/OperationStatus_Get.json
 */
async function workflowsGet() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const locationName = "westus";
  const workflowId = "828219ea-083e-48b5-89ea-8fd9991b2e75";
  const operationId = "14b50e24-f68d-4b29-a882-38be9dfb8bd1";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.operationStatusOperations.get(
    resourceGroupName,
    locationName,
    workflowId,
    operationId
  );
  console.log(result);
}

workflowsGet().catch(console.error);
```
