const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements an IP Community PUT method.
 *
 * @summary Implements an IP Community PUT method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpCommunities_Create_MaximumSet_Gen.json
 */
async function ipCommunitiesCreateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "rgIpCommunityLists";
  const ipCommunityName = "example-ipCommunity";
  const body = {
    action: "Permit",
    annotation: "annotationValue",
    communityMembers: ["1234:5678"],
    location: "EastUS",
    tags: { key2814: "" },
    wellKnownCommunities: ["Internet", "LocalAS", "NoExport", "GShut"],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipCommunities.beginCreateAndWait(
    resourceGroupName,
    ipCommunityName,
    body
  );
  console.log(result);
}
