const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists cryptographic certificate analysis results found in a firmware.
 *
 * @summary Lists cryptographic certificate analysis results found in a firmware.
 * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/CryptoCertificates_ListByFirmware_MaximumSet_Gen.json
 */
async function cryptoCertificatesListByFirmwareMaximumSetGen() {
  const subscriptionId =
    process.env["IOTFIRMWAREDEFENSE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["IOTFIRMWAREDEFENSE_RESOURCE_GROUP"] || "FirmwareAnalysisRG";
  const workspaceName = "default";
  const firmwareId = "109a9886-50bf-85a8-9d75-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cryptoCertificates.listByFirmware(
    resourceGroupName,
    workspaceName,
    firmwareId,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
