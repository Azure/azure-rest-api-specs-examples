const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements Network Fabric SKU GET method.
 *
 * @summary Implements Network Fabric SKU GET method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkFabricSkus_Get_MaximumSet_Gen.json
 */
async function networkFabricSkusGetMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const networkFabricSkuName = "example-fabricsku";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkFabricSkus.get(networkFabricSkuName);
  console.log(result);
}
