const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entity state (Etag) version of the Global policy definition in the Api Management service.
 *
 * @summary Gets the entity state (Etag) version of the Global policy definition in the Api Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadPolicy.json
 */
async function apiManagementHeadPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const policyId = "policy";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.policy.getEntityTag(resourceGroupName, serviceName, policyId);
  console.log(result);
}

apiManagementHeadPolicy().catch(console.error);
