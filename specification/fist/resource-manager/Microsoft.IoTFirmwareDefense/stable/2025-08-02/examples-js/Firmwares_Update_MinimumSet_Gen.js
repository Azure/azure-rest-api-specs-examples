const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to update firmware.
 *
 * @summary the operation to update firmware.
 * x-ms-original-file: 2025-08-02/Firmwares_Update_MinimumSet_Gen.json
 */
async function firmwaresUpdateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.firmwares.update("rgworkspaces-firmwares", "A7", "umrkdttp", {});
  console.log(result);
}
