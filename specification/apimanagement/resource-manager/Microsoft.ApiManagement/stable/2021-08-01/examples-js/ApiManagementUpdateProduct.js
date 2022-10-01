const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update existing product details.
 *
 * @summary Update existing product details.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateProduct.json
 */
async function apiManagementUpdateProduct() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "testproduct";
  const ifMatch = "*";
  const parameters = {
    displayName: "Test Template ProductName 4",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.product.update(
    resourceGroupName,
    serviceName,
    productId,
    ifMatch,
    parameters
  );
  console.log(result);
}

apiManagementUpdateProduct().catch(console.error);
