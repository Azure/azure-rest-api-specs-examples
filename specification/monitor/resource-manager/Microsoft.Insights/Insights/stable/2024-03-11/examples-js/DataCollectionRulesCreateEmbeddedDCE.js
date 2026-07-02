const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a data collection rule.
 *
 * @summary creates or updates a data collection rule.
 * x-ms-original-file: 2024-03-11/DataCollectionRulesCreateEmbeddedDCE.json
 */
async function createOrUpdateDataCollectionRuleWithEmbeddedIngestionEndpoints() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionRules.create("myResourceGroup", "myCollectionRule", {
    body: {
      kind: " Direct",
      location: "eastus",
      description: "A Direct Ingestion Rule with builtin ingestion fqdns",
      dataFlows: [
        {
          destinations: ["myworkspace"],
          outputStream: "Custom-LOGS1_CL",
          streams: ["Custom-LOGS1_CL"],
          transformKql:
            "source | extend jsonContext = parse_json(AdditionalContext) | project TimeGenerated = Time, Computer, AdditionalContext = jsonContext, CounterName=tostring(jsonContext.CounterName), CounterValue=toreal(jsonContext.CounterValue)",
        },
      ],
      destinations: {
        logAnalytics: [
          {
            name: "centralWorkspace",
            workspaceResourceId:
              "/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace",
          },
        ],
      },
      streamDeclarations: {
        "Custom-LOGS1_CL": {
          columns: [
            { name: "Time", type: "datetime" },
            { name: "Computer", type: "string" },
            { name: "AdditionalContext", type: "string" },
            { name: "CounterName", type: "string" },
            { name: "CounterValue", type: "real" },
          ],
        },
      },
    },
  });
  console.log(result);
}
