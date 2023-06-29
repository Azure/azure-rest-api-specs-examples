const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Displays L2IsolationDomains list by subscription GET method.
 *
 * @summary Displays L2IsolationDomains list by subscription GET method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/L2IsolationDomains_ListBySubscription_MaximumSet_Gen.json
 */
async function l2IsolationDomainsListBySubscriptionMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.l2IsolationDomains.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
