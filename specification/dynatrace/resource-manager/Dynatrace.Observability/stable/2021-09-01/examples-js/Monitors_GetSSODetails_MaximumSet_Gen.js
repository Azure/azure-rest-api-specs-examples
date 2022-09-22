const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the SSO configuration details from the partner.
 *
 * @summary Gets the SSO configuration details from the partner.
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_GetSSODetails_MaximumSet_Gen.json
 */
async function monitorsGetSsoDetailsMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const request = { userPrincipal: "alice@microsoft.com" };
  const options = { request };
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getSSODetails(resourceGroupName, monitorName, options);
  console.log(result);
}

monitorsGetSsoDetailsMaximumSetGen().catch(console.error);
