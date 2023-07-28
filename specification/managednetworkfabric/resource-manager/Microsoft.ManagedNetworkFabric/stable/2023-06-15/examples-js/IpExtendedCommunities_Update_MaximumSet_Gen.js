const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the IP Extended Community resource.
 *
 * @summary API to update certain properties of the IP Extended Community resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Update_MaximumSet_Gen.json
 */
async function ipExtendedCommunitiesUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const ipExtendedCommunityName = "example-ipExtendedCommunity";
  const body = {
    ipExtendedCommunityRules: [
      {
        action: "Permit",
        routeTargets: ["1234:2345"],
        sequenceNumber: 4155123341,
      },
    ],
    tags: { keyID: "KeyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.ipExtendedCommunities.beginUpdateAndWait(
    resourceGroupName,
    ipExtendedCommunityName,
    body
  );
  console.log(result);
}
