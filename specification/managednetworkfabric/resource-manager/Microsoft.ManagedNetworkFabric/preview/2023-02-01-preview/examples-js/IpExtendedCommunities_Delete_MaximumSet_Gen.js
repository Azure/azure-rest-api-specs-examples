const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements IP Extended Community DELETE method.
 *
 * @summary Implements IP Extended Community DELETE method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_Delete_MaximumSet_Gen.json
 */
async function ipExtendedCommunitiesDeleteMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "rgIpExtendedCommunityLists";
  const ipExtendedCommunityName = "example_ipExtendedCommunity";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipExtendedCommunities.beginDeleteAndWait(
    resourceGroupName,
    ipExtendedCommunityName
  );
  console.log(result);
}
