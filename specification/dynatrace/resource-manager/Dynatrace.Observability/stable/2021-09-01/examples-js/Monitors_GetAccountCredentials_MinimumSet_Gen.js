const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the user account credentials for a Monitor
 *
 * @summary Gets the user account credentials for a Monitor
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_GetAccountCredentials_MinimumSet_Gen.json
 */
async function monitorsGetAccountCredentialsMinimumSetGen() {
  const subscriptionId =
    process.env["DYNATRACE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DYNATRACE_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getAccountCredentials(resourceGroupName, monitorName);
  console.log(result);
}
