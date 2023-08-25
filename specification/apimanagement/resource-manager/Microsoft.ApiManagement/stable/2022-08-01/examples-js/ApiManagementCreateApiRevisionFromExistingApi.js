const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing specified API of the API Management service instance.
 *
 * @summary Creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiRevisionFromExistingApi.json
 */
async function apiManagementCreateApiRevisionFromExistingApi() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "echo-api;rev=3";
  const parameters = {
    path: "echo",
    apiRevisionDescription: "Creating a Revision of an existing API",
    serviceUrl: "http://echoapi.cloudapp.net/apiv3",
    sourceApiId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api",
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
