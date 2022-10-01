const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that API entity specified by identifier is associated with the Gateway entity.
 *
 * @summary Checks that API entity specified by identifier is associated with the Gateway entity.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadGatewayApi.json
 */
async function apiManagementHeadGatewayApi() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const apiId = "api1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gatewayApi.getEntityTag(
    resourceGroupName,
    serviceName,
    gatewayId,
    apiId
  );
  console.log(result);
}

apiManagementHeadGatewayApi().catch(console.error);
