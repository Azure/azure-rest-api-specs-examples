const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the specified product.
 *
 * @summary Returns the specified product.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/Product/GetPost.json
 */
async function returnsTheSpecifiedProduct() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const registrationName = "testregistration";
  const productName = "Microsoft.OSTCExtensions.VMAccessForLinux.1.4.7.1";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.products.getProduct(resourceGroup, registrationName, productName);
  console.log(result);
}

returnsTheSpecifiedProduct().catch(console.error);
