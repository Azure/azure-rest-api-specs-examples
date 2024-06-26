const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if Certificate entity is assigned to Gateway entity as Certificate Authority.
 *
 * @summary Checks if Certificate entity is assigned to Gateway entity as Certificate Authority.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadGatewayCertificateAuthority.json
 */
async function apiManagementHeadGatewayCertificateAuthority() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const certificateId = "cert1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gatewayCertificateAuthority.getEntityTag(
    resourceGroupName,
    serviceName,
    gatewayId,
    certificateId
  );
  console.log(result);
}
