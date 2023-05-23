const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new trunked network or update the properties of the existing trunked network.
 *
 * @summary Create a new trunked network or update the properties of the existing trunked network.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/TrunkedNetworks_Create.json
 */
async function createOrUpdateTrunkedNetwork() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const trunkedNetworkName = "trunkedNetworkName";
  const trunkedNetworkParameters = {
    extendedLocation: {
      name: "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName",
      type: "CustomLocation",
    },
    hybridAksPluginType: "DPDK",
    interfaceName: "eth0",
    isolationDomainIds: [
      "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l2IsolationDomains/l2IsolationDomainName",
      "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName",
    ],
    location: "location",
    tags: { key1: "myvalue1", key2: "myvalue2" },
    vlans: [12, 14],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.trunkedNetworks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    trunkedNetworkName,
    trunkedNetworkParameters
  );
  console.log(result);
}
