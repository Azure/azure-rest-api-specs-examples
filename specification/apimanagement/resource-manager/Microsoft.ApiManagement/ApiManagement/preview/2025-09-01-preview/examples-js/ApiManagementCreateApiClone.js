const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified API of the API Management service instance.
 *
 * @summary creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateApiClone.json
 */
async function apiManagementCreateApiClone() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.createOrUpdate("rg1", "apimService1", "echo-api2", {
    path: "echo2",
    description: "Copy of Existing Echo Api including Operations.",
    displayName: "Echo API2",
    isCurrent: true,
    protocols: ["http", "https"],
    serviceUrl: "http://echoapi.cloudapp.net/api",
    sourceApiId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/58a4aeac497000007d040001",
    subscriptionRequired: true,
  });
  console.log(result);
}
