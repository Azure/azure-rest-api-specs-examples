const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the active diagnostic settings for the specified resource.
 *
 * @summary Gets the active diagnostic settings for the specified resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/getDiagnosticSettingCategory.json
 */
async function getsTheDiagnosticSettingForCategory() {
  const resourceUri =
    "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
  const name = "mysetting";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.diagnosticSettings.get(resourceUri, name);
  console.log(result);
}
