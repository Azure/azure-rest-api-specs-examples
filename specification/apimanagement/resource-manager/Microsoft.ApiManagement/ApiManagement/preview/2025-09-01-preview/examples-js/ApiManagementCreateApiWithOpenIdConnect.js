const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified API of the API Management service instance.
 *
 * @summary creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateApiWithOpenIdConnect.json
 */
async function apiManagementCreateApiWithOpenIdConnect() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.createOrUpdate("rg1", "apimService1", "tempgroup", {
    path: "petstore",
    description:
      "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.",
    authenticationSettings: {
      openid: {
        bearerTokenSendingMethods: ["authorizationHeader"],
        openidProviderId: "testopenid",
      },
    },
    displayName: "Swagger Petstore",
    protocols: ["https"],
    serviceUrl: "http://petstore.swagger.io/v2",
    subscriptionKeyParameterNames: {
      header: "Ocp-Apim-Subscription-Key",
      query: "subscription-key",
    },
  });
  console.log(result);
}
