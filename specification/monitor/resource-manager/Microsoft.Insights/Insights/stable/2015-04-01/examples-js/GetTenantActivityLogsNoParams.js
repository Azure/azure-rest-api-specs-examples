const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the Activity Logs for the Tenant.<br>Everything that is applicable to the API to get the Activity Logs for the subscription is applicable to this API (the parameters, $filter, etc.).<br>One thing to point out here is that this API does *not* retrieve the logs at the individual subscription of the tenant but only surfaces the logs that were generated at the tenant level.
 *
 * @summary gets the Activity Logs for the Tenant.<br>Everything that is applicable to the API to get the Activity Logs for the subscription is applicable to this API (the parameters, $filter, etc.).<br>One thing to point out here is that this API does *not* retrieve the logs at the individual subscription of the tenant but only surfaces the logs that were generated at the tenant level.
 * x-ms-original-file: 2015-04-01/GetTenantActivityLogsNoParams.json
 */
async function getTenantActivityLogsWithoutFilterOrSelect() {
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (const item of client.tenantActivityLogs.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
