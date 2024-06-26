const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the log profile.
 *
 * @summary Gets the log profile.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getLogProfile.json
 */
async function getLogProfile() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "df602c9c-7aa0-407d-a6fb-eb20c8bd1192";
  const logProfileName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.logProfiles.get(logProfileName);
  console.log(result);
}
