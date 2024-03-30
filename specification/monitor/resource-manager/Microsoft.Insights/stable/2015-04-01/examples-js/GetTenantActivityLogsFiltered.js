const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Activity Logs for the Tenant.<br>Everything that is applicable to the API to get the Activity Logs for the subscription is applicable to this API (the parameters, $filter, etc.).<br>One thing to point out here is that this API does *not* retrieve the logs at the individual subscription of the tenant but only surfaces the logs that were generated at the tenant level.
 *
 * @summary Gets the Activity Logs for the Tenant.<br>Everything that is applicable to the API to get the Activity Logs for the subscription is applicable to this API (the parameters, $filter, etc.).<br>One thing to point out here is that this API does *not* retrieve the logs at the individual subscription of the tenant but only surfaces the logs that were generated at the tenant level.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetTenantActivityLogsFiltered.json
 */
async function getTenantActivityLogsWithFilter() {
  const filter =
    "eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (let item of client.tenantActivityLogs.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
