const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates diagnostic settings for the specified resource.
 *
 * @summary Creates or updates diagnostic settings for the specified resource.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/examples/createOrUpdateDiagnosticSettingCategory.json
 */
async function createsOrUpdatesTheDiagnosticSettingForCategory() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri =
    "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6";
  const name = "mysetting";
  const parameters = {
    eventHubAuthorizationRuleId:
      "/subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourceGroups/montest/providers/microsoft.eventhub/namespaces/mynamespace/authorizationrules/myrule",
    eventHubName: "myeventhub",
    logAnalyticsDestinationType: "Dedicated",
    logs: [
      {
        category: "WorkflowRuntime",
        enabled: true,
        retentionPolicy: { days: 0, enabled: false },
      },
    ],
    marketplacePartnerId:
      "/subscriptions/abcdeabc-1234-1234-ab12-123a1234567a/resourceGroups/test-rg/providers/Microsoft.Datadog/monitors/dd1",
    metrics: [
      {
        category: "WorkflowMetrics",
        enabled: true,
        retentionPolicy: { days: 0, enabled: false },
      },
    ],
    storageAccountId:
      "/subscriptions/df602c9c-7aa0-407d-a6fb-eb20c8bd1192/resourceGroups/apptest/providers/Microsoft.Storage/storageAccounts/appteststorage1",
    workspaceId: "",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.diagnosticSettings.createOrUpdate(resourceUri, name, parameters);
  console.log(result);
}

createsOrUpdatesTheDiagnosticSettingForCategory().catch(console.error);
