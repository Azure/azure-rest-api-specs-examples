const { IotHubClient } = require("@azure/arm-iothub-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Exports all the device identities in the IoT hub identity registry to an Azure Storage blob container. For more information, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-identity-registry#import-and-export-device-identities.
 *
 * @summary Exports all the device identities in the IoT hub identity registry to an Azure Storage blob container. For more information, see: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-identity-registry#import-and-export-device-identities.
 * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2019-07-01-preview/examples/iothub_exportdevices.json
 */
async function iotHubResourceExportDevices() {
  const subscriptionId =
    process.env["IOTHUB_SUBSCRIPTION_ID"] || "91d12660-3dec-467a-be2a-213b5544ddc0";
  const resourceGroupName = process.env["IOTHUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "testHub";
  const exportDevicesParameters = {
    excludeKeys: true,
    exportBlobContainerUri: "testBlob",
  };
  const credential = new DefaultAzureCredential();
  const client = new IotHubClient(credential, subscriptionId);
  const result = await client.iotHubResource.exportDevices(
    resourceGroupName,
    resourceName,
    exportDevicesParameters
  );
  console.log(result);
}
