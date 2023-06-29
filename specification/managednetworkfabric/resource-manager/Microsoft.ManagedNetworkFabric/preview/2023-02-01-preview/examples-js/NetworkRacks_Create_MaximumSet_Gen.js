const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Network Rack resource.
 *
 * @summary Create Network Rack resource.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkRacks_Create_MaximumSet_Gen.json
 */
async function networkRacksCreateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkRackName = "networkRackName";
  const body = {
    annotation: "null",
    location: "eastus",
    networkFabricId:
      "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkFabrics/networkFabricName",
    networkRackSku: "RackSKU",
    tags: { keyID: "keyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkRacks.beginCreateAndWait(
    resourceGroupName,
    networkRackName,
    body
  );
  console.log(result);
}
