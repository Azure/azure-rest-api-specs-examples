const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the firmware analysis workspaces in the specified subscription.
 *
 * @summary Lists all of the firmware analysis workspaces in the specified subscription.
 * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_ListBySubscription_MinimumSet_Gen.json
 */
async function workspacesListBySubscriptionMinimumSetGen() {
  const subscriptionId =
    process.env["IOTFIRMWAREDEFENSE_SUBSCRIPTION_ID"] || "5443A01A-5242-4950-AC1A-2DD362180254";
  const credential = new DefaultAzureCredential();
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.workspaces.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
