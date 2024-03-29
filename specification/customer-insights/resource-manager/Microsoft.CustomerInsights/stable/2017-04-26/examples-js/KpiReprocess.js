const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reprocesses the Kpi values of the specified KPI.
 *
 * @summary Reprocesses the Kpi values of the specified KPI.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/KpiReprocess.json
 */
async function kpiReprocess() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const kpiName = "kpiTest45453647";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.kpi.reprocess(resourceGroupName, hubName, kpiName);
  console.log(result);
}

kpiReprocess().catch(console.error);
