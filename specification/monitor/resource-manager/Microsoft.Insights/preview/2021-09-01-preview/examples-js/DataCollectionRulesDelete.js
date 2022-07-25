const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a data collection rule.
 *
 * @summary Deletes a data collection rule.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-09-01-preview/examples/DataCollectionRulesDelete.json
 */
async function deleteDataCollectionRule() {
  const subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
  const resourceGroupName = "myResourceGroup";
  const dataCollectionRuleName = "myCollectionRule";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.dataCollectionRules.delete(resourceGroupName, dataCollectionRuleName);
  console.log(result);
}

deleteDataCollectionRule().catch(console.error);
