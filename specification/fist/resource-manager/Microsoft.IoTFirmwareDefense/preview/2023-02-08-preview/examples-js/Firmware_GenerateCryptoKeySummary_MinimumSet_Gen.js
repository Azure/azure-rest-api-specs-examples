const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to provide a high level summary of the discovered cryptographic keys reported for the firmware image.
 *
 * @summary The operation to provide a high level summary of the discovered cryptographic keys reported for the firmware image.
 * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoKeySummary_MinimumSet_Gen.json
 */
async function firmwareGenerateCryptoKeySummaryMinimumSetGen() {
  const subscriptionId =
    process.env["IOTFIRMWAREDEFENSE_SUBSCRIPTION_ID"] || "C589E84A-5C11-4A25-9CF9-4E9C2F1EBFCA";
  const resourceGroupName =
    process.env["IOTFIRMWAREDEFENSE_RESOURCE_GROUP"] || "rgworkspaces-firmwares";
  const workspaceName = "j5QE_";
  const firmwareId = "wujtpcgypfpqseyrsebolarkspy";
  const credential = new DefaultAzureCredential();
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.firmwareOperations.generateCryptoKeySummary(
    resourceGroupName,
    workspaceName,
    firmwareId
  );
  console.log(result);
}
