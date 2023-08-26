const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates policy configuration for the API.
 *
 * @summary Creates or updates policy configuration for the API.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiPolicy.json
 */
async function apiManagementCreateApiPolicy() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "5600b57e7e8880006a040001";
  const policyId = "policy";
  const ifMatch = "*";
  const parameters = {
    format: "xml",
    value:
      "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>",
  };
  const options = { ifMatch };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiPolicy.createOrUpdate(
    resourceGroupName,
    serviceName,
    apiId,
    policyId,
    parameters,
    options
  );
  console.log(result);
}
