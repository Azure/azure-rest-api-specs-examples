const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a monitor resource.
 *
 * @summary Delete a monitor resource.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/Monitors_Delete.json
 */
async function monitorsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.monitors.beginDeleteAndWait(resourceGroupName, monitorName);
  console.log(result);
}

monitorsDelete().catch(console.error);
