const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves gateway keys.
 *
 * @summary Retrieves gateway keys.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGatewayListKeys.json
 */
async function apiManagementGatewayListKeys() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gateway.listKeys(resourceGroupName, serviceName, gatewayId);
  console.log(result);
}

apiManagementGatewayListKeys().catch(console.error);
