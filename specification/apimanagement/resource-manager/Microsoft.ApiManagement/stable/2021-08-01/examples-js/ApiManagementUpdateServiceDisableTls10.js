const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing API Management service.
 *
 * @summary Updates an existing API Management service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateServiceDisableTls10.json
 */
async function apiManagementUpdateServiceDisableTls10() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const parameters = {
    customProperties: {
      microsoftWindowsAzureApiManagementGatewaySecurityProtocolsTls10: "false",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    parameters
  );
  console.log(result);
}

apiManagementUpdateServiceDisableTls10().catch(console.error);
