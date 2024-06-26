const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the collection of subscriptions to the specified product.
 *
 * @summary Lists the collection of subscriptions to the specified product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductSubscriptions.json
 */
async function apiManagementListProductSubscriptions() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const productId = "5600b57e7e8880006a060002";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.productSubscriptions.list(
    resourceGroupName,
    serviceName,
    productId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
