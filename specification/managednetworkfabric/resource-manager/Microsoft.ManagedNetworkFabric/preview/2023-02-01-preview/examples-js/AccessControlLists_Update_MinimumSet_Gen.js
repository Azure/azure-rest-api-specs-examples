const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to API to update certain properties of the Access Control List resource.
 *
 * @summary API to update certain properties of the Access Control List resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/AccessControlLists_Update_MinimumSet_Gen.json
 */
async function accessControlListsUpdateMinimumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const accessControlListName = "aclOne";
  const body = {
    addressFamily: "ipv4",
    conditions: [
      {
        action: "allow",
        destinationAddress: "1.1.1.2",
        destinationPort: "21",
        sequenceNumber: 4,
        sourceAddress: "2.2.2.3",
        sourcePort: "65000",
        protocol: 6,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.accessControlLists.update(
    resourceGroupName,
    accessControlListName,
    body
  );
  console.log(result);
}
