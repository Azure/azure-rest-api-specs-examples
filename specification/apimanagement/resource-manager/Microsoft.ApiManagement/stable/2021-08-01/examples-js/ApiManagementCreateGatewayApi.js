const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds an API to the specified Gateway.
 *
 * @summary Adds an API to the specified Gateway.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGatewayApi.json
 */
async function apiManagementCreateGatewayApi() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const apiId = "echo-api";
  const parameters = { provisioningState: "created" };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gatewayApi.createOrUpdate(
    resourceGroupName,
    serviceName,
    gatewayId,
    apiId,
    options
  );
  console.log(result);
}

apiManagementCreateGatewayApi().catch(console.error);
