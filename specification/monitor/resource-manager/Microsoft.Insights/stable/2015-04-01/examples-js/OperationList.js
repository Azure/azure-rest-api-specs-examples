const { MonitorClient } = require("@azure/arm-monitor-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available operations from Microsoft.Insights provider.
 *
 * @summary Lists all of the available operations from Microsoft.Insights provider.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/OperationList.json
 */
async function getAListOfOperationsForAResourceProvider() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.operations.list();
  console.log(result);
}
