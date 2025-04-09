const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all of the features and SKUs offered by the Azure AI Search service in each region.
 *
 * @summary Lists all of the features and SKUs offered by the Azure AI Search service in each region.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2025-02-01-preview/examples/SearchListOfferings.json
 */
async function searchListOfferings() {
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.offerings.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
