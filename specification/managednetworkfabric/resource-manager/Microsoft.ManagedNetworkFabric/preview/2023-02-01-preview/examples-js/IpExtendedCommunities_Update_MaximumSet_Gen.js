const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the IP Extended Community resource.
 *
 * @summary API to update certain properties of the IP Extended Community resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_Update_MaximumSet_Gen.json
 */
async function ipExtendedCommunitiesUpdateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "rgIpExtendedCommunityLists";
  const ipExtendedCommunityName = "example_ipExtendedCommunity";
  const body = { tags: { key5054: "key1" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipExtendedCommunities.beginUpdateAndWait(
    resourceGroupName,
    ipExtendedCommunityName,
    body
  );
  console.log(result);
}
