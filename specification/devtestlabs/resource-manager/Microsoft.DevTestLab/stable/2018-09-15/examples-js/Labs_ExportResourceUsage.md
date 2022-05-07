Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Exports the lab resource usage into a storage account This operation can take a while to complete.
 *
 * @summary Exports the lab resource usage into a storage account This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ExportResourceUsage.json
 */
async function labsExportResourceUsage() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const exportResourceUsageParameters = {
    blobStorageAbsoluteSasUri:
      "https://invalid.blob.core.windows.net/export.blob?sv=2015-07-08&sig={sas}&sp=rcw",
    usageStartDate: new Date("2020-12-01T00:00:00Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.labs.beginExportResourceUsageAndWait(
    resourceGroupName,
    name,
    exportResourceUsageParameters
  );
  console.log(result);
}

labsExportResourceUsage().catch(console.error);
```
