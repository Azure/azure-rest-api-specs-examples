const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enable a receiver in an action group. This changes the receiver's status from Disabled to Enabled. This operation is only supported for Email or SMS receivers.
 *
 * @summary Enable a receiver in an action group. This changes the receiver's status from Disabled to Enabled. This operation is only supported for Email or SMS receivers.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-01/examples/enableReceiver.json
 */
async function enableTheReceiver() {
  const subscriptionId = "187f412d-1758-44d9-b052-169e2564721d";
  const resourceGroupName = "Default-NotificationRules";
  const actionGroupName = "SampleActionGroup";
  const enableRequest = { receiverName: "John Doe's mobile" };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.actionGroups.enableReceiver(
    resourceGroupName,
    actionGroupName,
    enableRequest
  );
  console.log(result);
}

enableTheReceiver().catch(console.error);
