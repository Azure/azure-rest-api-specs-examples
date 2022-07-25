const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes existing diagnostic settings for the specified resource.
 *
 * @summary Deletes existing diagnostic settings for the specified resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/deleteDiagnosticSetting.json
 */
async function deletesTheDiagnosticSetting() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
  const name = "mysetting";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.diagnosticSettings.delete(resourceUri, name);
  console.log(result);
}

deletesTheDiagnosticSetting().catch(console.error);
