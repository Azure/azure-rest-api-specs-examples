const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates new or updates existing specified API of the API Management service instance.
 *
 * @summary Creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApi.json
 */
async function apiManagementCreateApi() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const apiId = "tempgroup";
  const parameters = {
    path: "newapiPath",
    description: "apidescription5200",
    authenticationSettings: {
      oAuth2: {
        authorizationServerId: "authorizationServerId2283",
        scope: "oauth2scope2580",
      },
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
    parameters
  );
  console.log(result);
}
