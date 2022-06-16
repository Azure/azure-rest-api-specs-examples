const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Refresh the set password link and return a latest one.
 *
 * @summary Refresh the set password link and return a latest one.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/RefreshSetPassword_Get.json
 */
async function monitorsRefreshSetPasswordLink() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.monitors.refreshSetPasswordLink(resourceGroupName, monitorName);
  console.log(result);
}

monitorsRefreshSetPasswordLink().catch(console.error);
