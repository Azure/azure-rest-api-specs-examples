const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that API entity specified by identifier is associated with the Product entity.
 *
 * @summary Checks that API entity specified by identifier is associated with the Product entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadProductApi.json
 */
async function apiManagementHeadProductApi() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const productId = "5931a75ae4bbd512a88c680b";
  const apiId = "59306a29e4bbd510dc24e5f9";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.productApi.checkEntityExists(
    resourceGroupName,
    serviceName,
    productId,
    apiId
  );
  console.log(result);
}
