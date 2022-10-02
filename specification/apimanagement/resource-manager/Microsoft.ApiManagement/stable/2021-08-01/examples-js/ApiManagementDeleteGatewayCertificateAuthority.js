const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Remove relationship between Certificate Authority and Gateway entity.
 *
 * @summary Remove relationship between Certificate Authority and Gateway entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteGatewayCertificateAuthority.json
 */
async function apiManagementDeleteGatewayCertificateAuthority() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
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
    ifMatch
  );
  console.log(result);
}

apiManagementDeleteGatewayCertificateAuthority().catch(console.error);
