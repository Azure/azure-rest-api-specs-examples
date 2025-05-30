const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Assign Certificate entity to Gateway entity as Certificate Authority.
 *
 * @summary Assign Certificate entity to Gateway entity as Certificate Authority.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateGatewayCertificateAuthority.json
 */
async function apiManagementCreateGatewayCertificateAuthority() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const certificateId = "cert1";
  const parameters = { isTrusted: false };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gatewayCertificateAuthority.createOrUpdate(
    resourceGroupName,
    serviceName,
    gatewayId,
    certificateId,
    parameters,
  );
  console.log(result);
}
