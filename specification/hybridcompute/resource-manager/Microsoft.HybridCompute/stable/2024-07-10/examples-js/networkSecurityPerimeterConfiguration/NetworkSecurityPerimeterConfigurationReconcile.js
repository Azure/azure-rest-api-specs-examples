const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Forces the network security perimeter configuration to refresh for a private link scope.
 *
 * @summary Forces the network security perimeter configuration to refresh for a private link scope.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/stable/2024-07-10/examples/networkSecurityPerimeterConfiguration/NetworkSecurityPerimeterConfigurationReconcile.json
 */
async function reconcilesTheNetworkSecurityPerimeterConfigurationOfThePrivateLinkScope() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const perimeterName = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.myAssociation";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result =
    await client.networkSecurityPerimeterConfigurations.beginReconcileForPrivateLinkScopeAndWait(
      resourceGroupName,
      scopeName,
      perimeterName,
    );
  console.log(result);
}
