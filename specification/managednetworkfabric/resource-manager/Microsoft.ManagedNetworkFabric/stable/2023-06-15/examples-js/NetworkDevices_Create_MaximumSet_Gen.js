const { AzureNetworkFabricManagementServiceAPI } = require("@azure/arm-managednetworkfabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Network Device resource
 *
 * @summary Create a Network Device resource
 * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_Create_MaximumSet_Gen.json
 */
async function networkDevicesCreateMaximumSetGen() {
  const subscriptionId =
    process.env["MANAGEDNETWORKFABRIC_SUBSCRIPTION_ID"] || "1234ABCD-0A1B-1234-5678-123456ABCDEF";
  const resourceGroupName = process.env["MANAGEDNETWORKFABRIC_RESOURCE_GROUP"] || "example-rg";
  const networkDeviceName = "example-device";
  const body = {
    annotation: "annotation",
    hostName: "NFA-Device",
    location: "eastuseuap",
    networkDeviceSku: "DeviceSku",
    serialNumber: "Vendor;DCS-7280XXX-24;12.05;JPE2111XXXX",
    tags: { keyID: "KeyValue" },
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
