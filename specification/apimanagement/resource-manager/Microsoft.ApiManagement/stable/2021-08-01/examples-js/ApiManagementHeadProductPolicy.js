const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the ETag of the policy configuration at the Product level.
 *
 * @summary Get the ETag of the policy configuration at the Product level.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadProductPolicy.json
 */
async function apiManagementHeadProductPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "unlimited";
  const policyId = "policy";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.productPolicy.getEntityTag(
    resourceGroupName,
    serviceName,
    productId,
    policyId
  );
  console.log(result);
}

apiManagementHeadProductPolicy().catch(console.error);
