const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the IP Community resource.
 *
 * @summary API to update certain properties of the IP Community resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_Update_MaximumSet_Gen.json
 */
async function ipCommunitiesUpdateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const ipCommunityName = "example-ipCommunity";
  const body = { tags: { key2814: "" } };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipCommunities.beginUpdateAndWait(
    resourceGroupName,
    ipCommunityName,
    body
  );
  console.log(result);
}
