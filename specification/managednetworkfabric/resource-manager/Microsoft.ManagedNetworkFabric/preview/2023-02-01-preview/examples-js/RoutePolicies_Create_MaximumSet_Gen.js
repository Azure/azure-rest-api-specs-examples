const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements Route Policy PUT method.
 *
 * @summary Implements Route Policy PUT method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/RoutePolicies_Create_MaximumSet_Gen.json
 */
async function routePoliciesCreateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "rgRoutePolicies";
  const routePolicyName = "routePolicyName";
  const body = {
    annotation: "annotationValue",
    location: "EastUS",
    statements: [
      {
        action: {
          actionType: "Permit",
          ipCommunityProperties: {
            add: {
              ipCommunityIds: [
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName",
              ],
            },
            delete: {
              ipCommunityIds: [
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName",
              ],
            },
            set: {
              ipCommunityIds: [
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName",
              ],
            },
          },
          ipExtendedCommunityProperties: {
            add: {
              ipExtendedCommunityIds: [
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName",
              ],
            },
            delete: {
              ipExtendedCommunityIds: [
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName",
              ],
            },
            set: {
              ipExtendedCommunityIds: [
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName",
              ],
            },
          },
          localPreference: 20,
        },
        annotation: "annotationValue",
        condition: {
          ipCommunityIds: [
            "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipCommunities/ipCommunityName",
          ],
          ipExtendedCommunityIds: [
            "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/ipExtendedCommunityName",
          ],
          ipPrefixId:
            "subscriptions/xxxxxx/resourceGroups/resourcegroupname/providers/Microsoft.ManagedNetworkFabric/ipPrefixes/example-ipPrefix",
        },
        sequenceNumber: 7,
      },
    ],
    tags: { key8254: "" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.routePolicies.beginCreateAndWait(
    resourceGroupName,
    routePolicyName,
    body
  );
  console.log(result);
}
