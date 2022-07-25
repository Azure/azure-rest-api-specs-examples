const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available operations from Microsoft.Insights provider.
 *
 * @summary Lists all of the available operations from Microsoft.Insights provider.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/OperationList.json
 */
async function getAListOfOperationsForAResourceProvider() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.operations.list();
  console.log(result);
}

getAListOfOperationsForAResourceProvider().catch(console.error);
