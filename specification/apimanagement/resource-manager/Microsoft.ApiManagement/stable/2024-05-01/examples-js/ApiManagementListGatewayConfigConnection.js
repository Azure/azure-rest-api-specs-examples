const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all API Management gateway config connections within a gateway.
 *
 * @summary List all API Management gateway config connections within a gateway.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListGatewayConfigConnection.json
 */
async function apiManagementListGatewayConfigConnection() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const gatewayName = "standard-gw-1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.apiGatewayConfigConnection.listByGateway(
    resourceGroupName,
    gatewayName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
