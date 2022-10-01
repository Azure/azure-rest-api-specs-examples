const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a collection of the APIs associated with a product.
 *
 * @summary Lists a collection of the APIs associated with a product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductApis.json
 */
async function apiManagementListProductApis() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "5768181ea40f7eb6c49f6ac7";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.productApi.listByProduct(
    resourceGroupName,
    serviceName,
    productId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListProductApis().catch(console.error);
