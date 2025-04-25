const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates new or updates existing specified API of the API Management service instance.
 *
 * @summary Creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateApiWithMultipleOpenIdConnectProviders.json
 */
async function apiManagementCreateApiWithMultipleOpenIdConnectProviders() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "tempgroup";
  const parameters = {
    path: "newapiPath",
    description: "apidescription5200",
    authenticationSettings: {
      openidAuthenticationSettings: [
        {
          bearerTokenSendingMethods: ["authorizationHeader"],
          openidProviderId: "openidProviderId2283",
        },
        {
          bearerTokenSendingMethods: ["authorizationHeader"],
          openidProviderId: "openidProviderId2284",
        },
      ],
    },
    displayName: "apiname1463",
    protocols: ["https", "http"],
    serviceUrl: "http://newechoapi.cloudapp.net/api",
    subscriptionKeyParameterNames: { header: "header4520", query: "query3037" },
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
