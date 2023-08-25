const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing specified API of the API Management service instance.
 *
 * @summary Creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiUsingSwaggerImport.json
 */
async function apiManagementCreateApiUsingSwaggerImport() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "petstore";
  const parameters = {
    format: "swagger-link-json",
    path: "petstore",
    value: "http://petstore.swagger.io/v2/swagger.json",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    apiId,
    parameters
  );
  console.log(result);
}
