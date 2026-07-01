const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a data collection rule.
 *
 * @summary creates or updates a data collection rule.
 * x-ms-original-file: 2024-03-11/DataCollectionRulesCreateAgentSettings.json
 */
async function createOrUpdateAnAgentSettingsConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionRules.create("myResourceGroup", "myCollectionRule", {
    body: {
      kind: "AgentSettings",
      location: "eastus",
      description: "An agent settings configuration",
      agentSettings: {
        logs: [
          { name: "MaxDiskQuotaInMB", value: "5000" },
          { name: "UseTimeReceivedForForwardedEvents", value: "1" },
          { name: "Tags", value: "Azure, Monitoring" },
        ],
      },
    },
  });
  console.log(result);
}
