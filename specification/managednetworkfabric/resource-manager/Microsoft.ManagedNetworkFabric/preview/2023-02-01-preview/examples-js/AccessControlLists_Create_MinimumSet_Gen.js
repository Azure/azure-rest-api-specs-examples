const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements Access Control List PUT method.
 *
 * @summary Implements Access Control List PUT method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/AccessControlLists_Create_MinimumSet_Gen.json
 */
async function accessControlListsCreateMinimumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const accessControlListName = "aclOne";
  const body = {
    addressFamily: "ipv4",
    conditions: [
      {
        action: "allow",
        destinationAddress: "1.1.1.1",
        destinationPort: "21",
        sequenceNumber: 3,
        sourceAddress: "2.2.2.2",
        sourcePort: "65000",
        protocol: 6,
      },
    ],
    location: "EastUs",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.accessControlLists.create(
    resourceGroupName,
    accessControlListName,
    body
  );
  console.log(result);
}
