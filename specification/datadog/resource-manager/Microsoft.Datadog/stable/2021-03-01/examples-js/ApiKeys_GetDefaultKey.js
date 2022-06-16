const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the default api key.
 *
 * @summary Get the default api key.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/ApiKeys_GetDefaultKey.json
 */
async function monitorsGetDefaultKey() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.monitors.getDefaultKey(resourceGroupName, monitorName);
  console.log(result);
}

monitorsGetDefaultKey().catch(console.error);
