const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Delete product.
 *
 * @summary Delete product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteWorkspaceProduct.json
 */
async function apiManagementDeleteWorkspaceProduct() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const productId = "testproduct";
  const ifMatch = "*";
  const deleteSubscriptions = true;
  const options = { deleteSubscriptions };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceProduct.delete(
    resourceGroupName,
    serviceName,
    workspaceId,
    productId,
    ifMatch,
    options,
  );
  console.log(result);
}
