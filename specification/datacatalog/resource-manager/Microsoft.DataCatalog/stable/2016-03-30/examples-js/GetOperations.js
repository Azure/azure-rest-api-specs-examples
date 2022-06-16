const { DataCatalogRestClient } = require("@azure/arm-datacatalog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available Azure Data Catalog service operations.
 *
 * @summary Lists all the available Azure Data Catalog service operations.
 * x-ms-original-file: specification/datacatalog/resource-manager/Microsoft.DataCatalog/stable/2016-03-30/examples/GetOperations.json
 */
async function getOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new DataCatalogRestClient(credential, subscriptionId);
  const result = await client.aDCOperations.list();
  console.log(result);
}

getOperations().catch(console.error);
