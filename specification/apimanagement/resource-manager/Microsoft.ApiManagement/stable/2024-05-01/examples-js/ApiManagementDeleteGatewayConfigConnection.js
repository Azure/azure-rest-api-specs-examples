const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes an existing API Management gateway config connection.
 *
 * @summary Deletes an existing API Management gateway config connection.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteGatewayConfigConnection.json
 */
async function apiManagementGatewayDeleteGateway() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const gatewayName = "standard-gw-01";
  const configConnectionName = "gcc-01";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiGatewayConfigConnection.beginDeleteAndWait(
    resourceGroupName,
    gatewayName,
    configConnectionName,
    ifMatch,
  );
  console.log(result);
}
