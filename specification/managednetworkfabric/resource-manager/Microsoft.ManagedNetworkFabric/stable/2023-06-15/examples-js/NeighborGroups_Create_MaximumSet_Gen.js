const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements the Neighbor Group PUT method.
 *
 * @summary Implements the Neighbor Group PUT method.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NeighborGroups_Create_MaximumSet_Gen.json
 */
async function neighborGroupsCreateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const neighborGroupName = "example-neighborGroup";
  const body = {
    annotation: "annotation",
    destination: {
      ipv4Addresses: [
        "10.10.10.10",
        "20.10.10.10",
        "30.10.10.10",
        "40.10.10.10",
        "50.10.10.10",
        "60.10.10.10",
        "70.10.10.10",
        "80.10.10.10",
        "90.10.10.10",
      ],
      ipv6Addresses: ["2F::/100"],
    },
    location: "eastus",
    tags: { key8107: "1234" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.neighborGroups.beginCreateAndWait(
    resourceGroupName,
    neighborGroupName,
    body
  );
  console.log(result);
}
