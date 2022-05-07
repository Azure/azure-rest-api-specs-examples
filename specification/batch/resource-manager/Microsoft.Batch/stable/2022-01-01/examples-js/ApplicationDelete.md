Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-batch_7.1.1/sdk/batch/arm-batch/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an application.
 *
 * @summary Deletes an application.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationDelete.json
 */
async function applicationDelete() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const applicationName = "app1";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.applicationOperations.delete(
    resourceGroupName,
    accountName,
    applicationName
  );
  console.log(result);
}

applicationDelete().catch(console.error);
```
