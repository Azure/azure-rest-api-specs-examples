const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Shared Access Authorization Token for the gateway.
 *
 * @summary Gets the Shared Access Authorization Token for the gateway.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayGenerateToken.json
 */
async function apiManagementGatewayGenerateToken() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const parameters = {
    expiry: new Date("2020-04-21T00:44:24.2845269Z"),
    keyType: "primary",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gateway.generateToken(
    resourceGroupName,
    serviceName,
    gatewayId,
    parameters
  );
  console.log(result);
}

apiManagementGatewayGenerateToken().catch(console.error);
