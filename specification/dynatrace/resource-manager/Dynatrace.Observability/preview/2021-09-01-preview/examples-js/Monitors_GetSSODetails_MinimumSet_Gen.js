const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the SSO configuration details from the partner.
 *
 * @summary Gets the SSO configuration details from the partner.
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/Monitors_GetSSODetails_MinimumSet_Gen.json
 */
async function monitorsGetSsoDetailsMinimumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const request = {};
  const options = { request };
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getSSODetails(resourceGroupName, monitorName, options);
  console.log(result);
}

monitorsGetSsoDetailsMinimumSetGen().catch(console.error);
