const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all Network Rack SKUs in the given subscription.
 *
 * @summary List all Network Rack SKUs in the given subscription.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkRackSkus_ListBySubscription_MaximumSet_Gen.json
 */
async function networkRackSkusListBySubscriptionMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkRackSkus.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
