const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the policy configuration at the Api Operation.
 *
 * @summary Deletes the policy configuration at the Api Operation.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiOperationPolicy.json
 */
async function apiManagementDeleteApiOperationPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "testapi";
  const operationId = "testoperation";
  const policyId = "policy";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiOperationPolicy.delete(
    resourceGroupName,
    serviceName,
    apiId,
    operationId,
    policyId,
    ifMatch
  );
  console.log(result);
}

apiManagementDeleteApiOperationPolicy().catch(console.error);
