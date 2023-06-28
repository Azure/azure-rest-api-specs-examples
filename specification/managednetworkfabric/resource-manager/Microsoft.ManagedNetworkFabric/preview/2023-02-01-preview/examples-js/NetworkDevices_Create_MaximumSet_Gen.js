const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Network Device resource
 *
 * @summary Create a Network Device resource
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/NetworkDevices_Create_MaximumSet_Gen.json
 */
async function networkDevicesCreateMaximumSetGen() {
  const subscriptionId = process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName =
    process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "resourceGroupName";
  const networkDeviceName = "networkDeviceName";
  const body = {
    annotation: "null",
    hostName: "networkDeviceName",
    location: "eastus",
    networkDeviceRole: "CE",
    networkDeviceSku: "DefaultSku",
    serialNumber: "Arista;DCS-7280PR3-24;12.05;JPE21330382",
    tags: { keyID: "keyValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureNetworkFabricManagementServiceAPI(credential, subscriptionId);
  const result = await client.networkDevices.beginCreateAndWait(
    resourceGroupName,
    networkDeviceName,
    body
  );
  console.log(result);
}
