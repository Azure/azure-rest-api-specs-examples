const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the static interface maps for the given Network Device.
 *
 * @summary Get the static interface maps for the given Network Device.
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDevices_getStaticInterfaceMaps_MaximumSet_Gen.json
 */
async function networkDevicesGetStaticInterfaceMapsMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkDeviceName = "networkDeviceName";
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkDevices.beginGetStaticInterfaceMapsAndWait(
    resourceGroupName,
    networkDeviceName
  );
  console.log(result);
}
