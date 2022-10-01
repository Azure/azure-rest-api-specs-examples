const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete product.
 *
 * @summary Delete product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteProduct.json
 */
async function apiManagementDeleteProduct() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "testproduct";
  const ifMatch = "*";
  const deleteSubscriptions = true;
  const options = { deleteSubscriptions };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.product.delete(
    resourceGroupName,
    serviceName,
    productId,
    ifMatch,
    options
  );
  console.log(result);
}

apiManagementDeleteProduct().catch(console.error);
