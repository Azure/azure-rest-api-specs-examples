const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Regenerates specified gateway key invalidating any tokens created with it.
 *
 * @summary Regenerates specified gateway key invalidating any tokens created with it.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGatewayRegenerateKey.json
 */
async function apiManagementGatewayRegenerateKey() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gwId";
  const parameters = {
    keyType: "primary",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gateway.regenerateKey(
    resourceGroupName,
    serviceName,
    gatewayId,
    parameters,
  );
  console.log(result);
}
