const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified API of the API Management service instance.
 *
 * @summary creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateODataApi.json
 */
async function apiManagementCreateODataApi() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.createOrUpdate("rg1", "apimService1", "tempgroup", {
    apiType: "odata",
    format: "odata-link",
    path: "odata-api",
    description: "apidescription5200",
    displayName: "apiname1463",
    protocols: ["http", "https"],
    serviceUrl: "https://services.odata.org/TripPinWebApiService",
    value: "https://services.odata.org/TripPinWebApiService/$metadata",
  });
  console.log(result);
}
