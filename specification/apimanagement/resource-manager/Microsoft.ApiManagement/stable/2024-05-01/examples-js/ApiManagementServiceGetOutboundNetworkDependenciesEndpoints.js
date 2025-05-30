const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the network endpoints of all outbound dependencies of a ApiManagement service.
 *
 * @summary Gets the network endpoints of all outbound dependencies of a ApiManagement service.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementServiceGetOutboundNetworkDependenciesEndpoints.json
 */
async function apiManagementServiceGetOutboundNetworkDependenciesEndpoints() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.outboundNetworkDependenciesEndpoints.listByService(
    resourceGroupName,
    serviceName,
  );
  console.log(result);
}
