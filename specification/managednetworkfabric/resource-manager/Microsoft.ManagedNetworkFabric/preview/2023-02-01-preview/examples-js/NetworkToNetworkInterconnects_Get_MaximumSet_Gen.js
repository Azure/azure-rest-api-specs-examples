const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements NetworkToNetworkInterconnects GET method.
 *
 * @summary Implements NetworkToNetworkInterconnects GET method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkToNetworkInterconnects_Get_MaximumSet_Gen.json
 */
async function networkToNetworkInterconnectsGetMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkFabricName = "FabricName";
  const networkToNetworkInterconnectName = "DefaultNNI";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkToNetworkInterconnects.get(
    resourceGroupName,
    networkFabricName,
    networkToNetworkInterconnectName
  );
  console.log(result);
}
