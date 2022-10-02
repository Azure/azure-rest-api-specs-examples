const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the collection of developer groups associated with the specified product.
 *
 * @summary Lists the collection of developer groups associated with the specified product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductGroups.json
 */
async function apiManagementListProductGroups() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "5600b57e7e8880006a060002";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.productGroup.listByProduct(
    resourceGroupName,
    serviceName,
    productId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListProductGroups().catch(console.error);
