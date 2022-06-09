```javascript
const { DataCatalogRestClient } = require("@azure/arm-datacatalog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The List catalogs in Resource Group operation lists all the Azure Data Catalogs available under the given resource group.
 *
 * @summary The List catalogs in Resource Group operation lists all the Azure Data Catalogs available under the given resource group.
 * x-ms-original-file: specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/ListADCCatalogsByResourceGroup.json
 */
async function listAzureDataCatalogService() {
  const subscriptionId = "12345678-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new DataCatalogRestClient(credential, subscriptionId);
  const result = await client.aDCCatalogs.listtByResourceGroup(resourceGroupName);
  console.log(result);
}

listAzureDataCatalogService().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datacatalog_4.0.0/sdk/datacatalog/arm-datacatalog/README.md) on how to add the SDK to your project and authenticate.
