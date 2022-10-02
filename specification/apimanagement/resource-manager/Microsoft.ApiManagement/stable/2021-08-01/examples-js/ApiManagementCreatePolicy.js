const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the global policy configuration of the Api Management service.
 *
 * @summary Creates or updates the global policy configuration of the Api Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreatePolicy.json
 */
async function apiManagementCreatePolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const policyId = "policy";
  const parameters = {
    format: "xml",
    value:
      "<policies>  <inbound />  <backend>    <forward-request />  </backend>  <outbound /></policies>",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.policy.createOrUpdate(
    resourceGroupName,
    serviceName,
    policyId,
    parameters
  );
  console.log(result);
}

apiManagementCreatePolicy().catch(console.error);
