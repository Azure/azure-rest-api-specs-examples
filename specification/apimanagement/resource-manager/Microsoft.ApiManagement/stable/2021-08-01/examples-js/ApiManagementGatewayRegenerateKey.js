const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates specified gateway key invalidating any tokens created with it.
 *
 * @summary Regenerates specified gateway key invalidating any tokens created with it.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayRegenerateKey.json
 */
async function apiManagementGatewayRegenerateKey() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
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
    parameters
  );
  console.log(result);
}

apiManagementGatewayRegenerateKey().catch(console.error);
