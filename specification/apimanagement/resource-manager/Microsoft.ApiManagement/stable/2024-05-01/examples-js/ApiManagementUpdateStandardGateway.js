const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates an existing API Management gateway.
 *
 * @summary Updates an existing API Management gateway.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateStandardGateway.json
 */
async function apiManagementUpdateStandardGateway() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const gatewayName = "apimGateway1";
  const parameters = {
    sku: { name: "Standard", capacity: 10 },
    tags: { name: "Contoso", test: "User" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiGateway.beginUpdateAndWait(
    resourceGroupName,
    gatewayName,
    parameters,
  );
  console.log(result);
}
