const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get firmware.
 *
 * @summary get firmware.
 * x-ms-original-file: 2025-08-02/Firmwares_Get_MinimumSet_Gen.json
 */
async function firmwaresGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.firmwares.get("rgworkspaces-firmwares", "A7", "umrkdttp");
  console.log(result);
}
