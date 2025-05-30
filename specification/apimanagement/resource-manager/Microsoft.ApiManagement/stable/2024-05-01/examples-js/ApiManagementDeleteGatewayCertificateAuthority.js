const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Remove relationship between Certificate Authority and Gateway entity.
 *
 * @summary Remove relationship between Certificate Authority and Gateway entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeleteGatewayCertificateAuthority.json
 */
async function apiManagementDeleteGatewayCertificateAuthority() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const certificateId = "default";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gatewayCertificateAuthority.delete(
    resourceGroupName,
    serviceName,
    gatewayId,
    certificateId,
    ifMatch,
  );
  console.log(result);
}
