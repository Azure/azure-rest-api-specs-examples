const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements IpCommunities list by subscription GET method.
 *
 * @summary Implements IpCommunities list by subscription GET method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_ListBySubscription_MaximumSet_Gen.json
 */
async function ipCommunitiesListBySubscriptionMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.ipCommunities.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
