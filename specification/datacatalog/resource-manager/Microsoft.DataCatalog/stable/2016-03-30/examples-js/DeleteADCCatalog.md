Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datacatalog_4.0.0/sdk/datacatalog/arm-datacatalog/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataCatalogRestClient } = require("@azure/arm-datacatalog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Delete Azure Data Catalog Service operation deletes an existing data catalog.
 *
 * @summary The Delete Azure Data Catalog Service operation deletes an existing data catalog.
 * x-ms-original-file: specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/DeleteADCCatalog.json
 */
async function deleteAzureDataCatalogService() {
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const catalogName = "exampleCatalog";
  const credential = new DefaultAzureCredential();
  const client = new DataCatalogRestClient(credential, subscriptionId);
  const result = await client.aDCCatalogs.beginDeleteAndWait(resourceGroupName, catalogName);
  console.log(result);
}

deleteAzureDataCatalogService().catch(console.error);
```
