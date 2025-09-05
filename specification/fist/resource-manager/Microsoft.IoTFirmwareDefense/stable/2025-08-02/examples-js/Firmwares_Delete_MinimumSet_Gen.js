const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete a firmware.
 *
 * @summary the operation to delete a firmware.
 * x-ms-original-file: 2025-08-02/Firmwares_Delete_MinimumSet_Gen.json
 */
async function firmwaresDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  await client.firmwares.delete("rgworkspaces-firmwares", "A7", "umrkdttp");
}
