const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists the network security perimeter configurations for a private link scope.
 *
 * @summary Lists the network security perimeter configurations for a private link scope.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationList.json
 */
async function getsTheListOfNetworkSecurityPerimeterConfigurationsOfThePrivateLinkScope() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.networkSecurityPerimeterConfigurations.listByPrivateLinkScope(
    resourceGroupName,
    scopeName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
