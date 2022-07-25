const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the diagnostic settings category for the specified resource.
 *
 * @summary Gets the diagnostic settings category for the specified resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2017-05-01-preview/examples/getDiagnosticSettingsCategory.json
 */
async function getsTheDiagnosticSetting() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
  const name = "WorkflowRuntime";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.diagnosticSettingsCategory.get(resourceUri, name);
  console.log(result);
}

getsTheDiagnosticSetting().catch(console.error);
