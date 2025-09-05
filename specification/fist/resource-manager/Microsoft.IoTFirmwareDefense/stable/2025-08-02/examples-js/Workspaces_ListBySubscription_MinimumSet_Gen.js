const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all of the firmware analysis workspaces in the specified subscription.
 *
 * @summary lists all of the firmware analysis workspaces in the specified subscription.
 * x-ms-original-file: 2025-08-02/Workspaces_ListBySubscription_MinimumSet_Gen.json
 */
async function workspacesListBySubscriptionMaximumSetGenGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.workspaces.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
