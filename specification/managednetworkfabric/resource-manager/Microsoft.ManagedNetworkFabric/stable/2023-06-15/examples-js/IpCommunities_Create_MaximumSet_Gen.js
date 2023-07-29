const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements an IP Community PUT method.
 *
 * @summary Implements an IP Community PUT method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpCommunities_Create_MaximumSet_Gen.json
 */
async function ipCommunitiesCreateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const ipCommunityName = "example-ipcommunity";
  const body = {
    annotation: "annotation",
    ipCommunityRules: [
      {
        action: "Permit",
        communityMembers: ["1:1"],
        sequenceNumber: 4155123341,
        wellKnownCommunities: ["Internet"],
      },
    ],
    location: "eastus",
    tags: { keyId: "KeyValue" },
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
