const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements NetworkToNetworkInterconnects GET method.
 *
 * @summary Implements NetworkToNetworkInterconnects GET method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkToNetworkInterconnects_Get_MaximumSet_Gen.json
 */
async function networkToNetworkInterconnectsGetMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const networkFabricName = "example-fabric";
  const networkToNetworkInterconnectName = "example-nni";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkToNetworkInterconnects.get(
    resourceGroupName,
    networkFabricName,
    networkToNetworkInterconnectName
  );
  console.log(result);
}
