const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Adds the association between the specified developer group with the specified product.
 *
 * @summary Adds the association between the specified developer group with the specified product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateProductGroup.json
 */
async function apiManagementCreateProductGroup() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const productId = "testproduct";
  const groupId = "templateGroup";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.productGroup.createOrUpdate(
    resourceGroupName,
    serviceName,
    productId,
    groupId,
  );
  console.log(result);
}
