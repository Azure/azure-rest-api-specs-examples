const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of products.
 *
 * @summary Returns a list of products.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/Product/ListPost.json
 */
async function returnsAListOfProducts() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const registrationName = "testregistration";
  const productName = "_all";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.products.getProducts(resourceGroup, registrationName, productName);
  console.log(result);
}

returnsAListOfProducts().catch(console.error);
