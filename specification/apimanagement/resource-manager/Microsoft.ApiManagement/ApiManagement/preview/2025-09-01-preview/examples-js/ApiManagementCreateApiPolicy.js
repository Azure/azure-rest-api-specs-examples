const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates policy configuration for the API.
 *
 * @summary creates or updates policy configuration for the API.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateApiPolicy.json
 */
async function apiManagementCreateApiPolicy() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiPolicy.createOrUpdate(
    "rg1",
    "apimService1",
    "5600b57e7e8880006a040001",
    "policy",
    {
      format: "xml",
      value:
        "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>",
    },
    { ifMatch: "*" },
  );
  console.log(result);
}
