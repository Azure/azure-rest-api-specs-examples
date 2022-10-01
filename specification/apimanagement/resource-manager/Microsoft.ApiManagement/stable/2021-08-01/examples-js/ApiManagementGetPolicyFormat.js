const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Global policy definition of the Api Management service.
 *
 * @summary Get the Global policy definition of the Api Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetPolicyFormat.json
 */
async function apiManagementGetPolicyFormat() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const policyId = "policy";
  const format = "rawxml";
  const options = { format };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.policy.get(resourceGroupName, serviceName, policyId, options);
  console.log(result);
}

apiManagementGetPolicyFormat().catch(console.error);
