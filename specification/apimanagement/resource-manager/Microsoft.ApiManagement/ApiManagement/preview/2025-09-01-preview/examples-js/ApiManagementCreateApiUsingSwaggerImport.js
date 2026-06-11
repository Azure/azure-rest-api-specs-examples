const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified API of the API Management service instance.
 *
 * @summary creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateApiUsingSwaggerImport.json
 */
async function apiManagementCreateApiUsingSwaggerImport() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.createOrUpdate("rg1", "apimService1", "petstore", {
    format: "swagger-link-json",
    path: "petstore",
    value: "http://petstore.swagger.io/v2/swagger.json",
  });
  console.log(result);
}
