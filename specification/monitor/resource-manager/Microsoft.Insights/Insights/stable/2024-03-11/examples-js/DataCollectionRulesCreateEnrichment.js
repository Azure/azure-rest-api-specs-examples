const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a data collection rule.
 *
 * @summary creates or updates a data collection rule.
 * x-ms-original-file: 2024-03-11/DataCollectionRulesCreateEnrichment.json
 */
async function createOrUpdateDataCollectionRuleWithEnrichment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionRules.create("myResourceGroup", "myCollectionRule", {
    body: {
      location: "eastus",
      description: "A rule showcasing ingestion time enrichment",
      dataCollectionEndpointId:
        "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myDataCollectionEndpoint",
      dataFlows: [
        {
          destinations: ["centralWorkspace"],
          outputStream: "Custom-LOGS1_CL",
          streams: ["Custom-TabularDataABC"],
          transformKql:
            "source | extend LookupData = lookup_string_am('mytextdatastore', Message) | project TimeGenerated, Message, AdditionalContext = LookupData.Message",
        },
      ],
      dataSources: {
        logFiles: [
          {
            name: "myTabularLogDataSource",
            format: "text",
            filePatterns: ["C:\\JavaLogs\\*\\*.log"],
            settings: { text: { recordStartTimestampFormat: "ISO 8601" } },
            streams: ["Custom-TabularDataABC"],
          },
        ],
      },
      destinations: {
        logAnalytics: [
          {
            name: "centralWorkspace",
            workspaceResourceId:
              "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace",
          },
        ],
      },
      references: {
        enrichmentData: {
          storageBlobs: [
            {
              name: "mytextdatastore",
              blobUrl: "https://myenrichmentstorage.blob.core.windows.net/enrichment",
              lookupType: "String",
              resourceId:
                "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourcegroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myenrichmentstorage",
            },
          ],
        },
      },
      streamDeclarations: {
        "Custom-TabularDataABC": {
          columns: [
            { name: "TimeGenerated", type: "datetime" },
            { name: "Message", type: "string" },
            { name: "AdditionalContext", type: "string" },
          ],
        },
      },
    },
  });
  console.log(result);
}
