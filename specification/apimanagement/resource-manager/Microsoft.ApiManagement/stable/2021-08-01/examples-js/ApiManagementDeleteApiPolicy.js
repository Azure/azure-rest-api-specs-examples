const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the policy configuration at the Api.
 *
 * @summary Deletes the policy configuration at the Api.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiPolicy.json
 */
async function apiManagementDeleteApiPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "loggerId";
  const policyId = "policy";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiPolicy.delete(
    resourceGroupName,
    serviceName,
    apiId,
    policyId,
    ifMatch
  );
  console.log(result);
}

apiManagementDeleteApiPolicy().catch(console.error);
