const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Gateway to be used in Api Management instance.
 *
 * @summary Creates or updates a Gateway to be used in Api Management instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateGateway.json
 */
async function apiManagementCreateGateway() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const gatewayId = "gw1";
  const parameters = {
    description: "my gateway 1",
    locationData: { name: "my location" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.gateway.createOrUpdate(
    resourceGroupName,
    serviceName,
    gatewayId,
    parameters
  );
  console.log(result);
}

apiManagementCreateGateway().catch(console.error);
