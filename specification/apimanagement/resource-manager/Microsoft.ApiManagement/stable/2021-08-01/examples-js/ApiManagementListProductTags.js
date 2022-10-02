const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Tags associated with the Product.
 *
 * @summary Lists all Tags associated with the Product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductTags.json
 */
async function apiManagementListProductTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "57d2ef278aa04f0888cba3f1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tag.listByProduct(resourceGroupName, serviceName, productId)) {
    resArray.push(item);
  }
  console.log(resArray);
}

apiManagementListProductTags().catch(console.error);
