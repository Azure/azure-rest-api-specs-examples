const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Detach the tag from the Product.
 *
 * @summary Detach the tag from the Product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteProductTag.json
 */
async function apiManagementDeleteProductTag() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const productId = "59d5b28d1f7fab116c282650";
  const tagId = "59d5b28e1f7fab116402044e";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.tag.detachFromProduct(
    resourceGroupName,
    serviceName,
    productId,
    tagId
  );
  console.log(result);
}
