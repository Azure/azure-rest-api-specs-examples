const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the list of records from the activity logs.
 *
 * @summary Provides the list of records from the activity logs.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2015-04-01/examples/GetActivityLogsFilteredAndSelected.json
 */
async function getActivityLogsWithFilterAndSelect() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "089bd33f-d4ec-47fe-8ba5-0753aa5c5b33";
  const filter =
    "eventTimestamp ge '2015-01-21T20:00:00Z' and eventTimestamp le '2015-01-23T20:00:00Z' and resourceGroupName eq 'MSSupportGroup'";
  const select =
    "eventName,id,resourceGroupName,resourceProviderName,operationName,status,eventTimestamp,correlationId,submissionTimestamp,level";
  const options = { select };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.activityLogs.list(filter, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
