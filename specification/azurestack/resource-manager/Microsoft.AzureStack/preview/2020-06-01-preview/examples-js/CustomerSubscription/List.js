const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of products.
 *
 * @summary Returns a list of products.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/CustomerSubscription/List.json
 */
async function returnsAListOfProducts() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const registrationName = "testregistration";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.customerSubscriptions.list(resourceGroup, registrationName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

returnsAListOfProducts().catch(console.error);
