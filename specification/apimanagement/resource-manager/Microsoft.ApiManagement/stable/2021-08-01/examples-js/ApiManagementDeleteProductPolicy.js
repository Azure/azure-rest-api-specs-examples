const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the policy configuration at the Product.
 *
 * @summary Deletes the policy configuration at the Product.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteProductPolicy.json
 */
async function apiManagementDeleteProductPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const productId = "testproduct";
  const policyId = "policy";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.productPolicy.delete(
    resourceGroupName,
    serviceName,
    productId,
    policyId,
    ifMatch
  );
  console.log(result);
}

apiManagementDeleteProductPolicy().catch(console.error);
