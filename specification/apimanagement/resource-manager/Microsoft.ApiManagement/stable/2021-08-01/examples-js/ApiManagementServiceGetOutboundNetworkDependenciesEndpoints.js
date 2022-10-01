const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the network endpoints of all outbound dependencies of a ApiManagement service.
 *
 * @summary Gets the network endpoints of all outbound dependencies of a ApiManagement service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementServiceGetOutboundNetworkDependenciesEndpoints.json
 */
async function apiManagementServiceGetOutboundNetworkDependenciesEndpoints() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.outboundNetworkDependenciesEndpoints.listByService(
    resourceGroupName,
    serviceName
  );
  console.log(result);
}

apiManagementServiceGetOutboundNetworkDependenciesEndpoints().catch(console.error);
