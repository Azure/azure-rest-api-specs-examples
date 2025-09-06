const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get an analysis result summary of a firmware by name.
 *
 * @summary get an analysis result summary of a firmware by name.
 * x-ms-original-file: 2025-08-02/Summaries_Get_MaximumSet_Gen.json
 */
async function summariesGetMaximumSetGenGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.summaries.get(
    "rgiotfirmwaredefense",
    "exampleWorkspaceName",
    "00000000-0000-0000-0000-000000000000",
    "Firmware",
  );
  console.log(result);
}
