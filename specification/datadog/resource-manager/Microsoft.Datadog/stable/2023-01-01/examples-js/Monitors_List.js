const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all monitors under the specified subscription.
 *
 * @summary List all monitors under the specified subscription.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/Monitors_List.json
 */
async function monitorsList() {
  const subscriptionId =
    process.env["DATADOG_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
