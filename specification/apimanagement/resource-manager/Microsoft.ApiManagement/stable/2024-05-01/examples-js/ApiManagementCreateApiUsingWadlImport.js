const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates new or updates existing specified API of the API Management service instance.
 *
 * @summary Creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateApiUsingWadlImport.json
 */
async function apiManagementCreateApiUsingWadlImport() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "petstore";
  const parameters = {
    format: "wadl-link-json",
    path: "collector",
    value:
      "https://developer.cisco.com/media/wae-release-6-2-api-reference/wae-collector-rest-api/application.wadl",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    apiId,
    parameters,
  );
  console.log(result);
}
