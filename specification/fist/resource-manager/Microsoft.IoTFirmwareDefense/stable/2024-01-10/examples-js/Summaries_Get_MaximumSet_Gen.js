const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get an analysis result summary of a firmware by name.
 *
 * @summary Get an analysis result summary of a firmware by name.
 * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Summaries_Get_MaximumSet_Gen.json
 */
async function summariesGetMaximumSetGen() {
  const subscriptionId =
    process.env["IOTFIRMWAREDEFENSE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["IOTFIRMWAREDEFENSE_RESOURCE_GROUP"] || "FirmwareAnalysisRG";
  const workspaceName = "default";
  const firmwareId = "109a9886-50bf-85a8-9d75-000000000000";
  const summaryName = "Firmware";
  const credential = new DefaultAzureCredential();
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.summaries.get(
    resourceGroupName,
    workspaceName,
    firmwareId,
    summaryName,
  );
  console.log(result);
}
