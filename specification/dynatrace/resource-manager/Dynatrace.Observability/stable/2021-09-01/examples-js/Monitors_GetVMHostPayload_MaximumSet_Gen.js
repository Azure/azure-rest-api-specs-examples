const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the payload that needs to be passed in the request body for installing Dynatrace agent on a VM.
 *
 * @summary Returns the payload that needs to be passed in the request body for installing Dynatrace agent on a VM.
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_GetVMHostPayload_MaximumSet_Gen.json
 */
async function monitorsGetVMHostPayloadMaximumSetGen() {
  const subscriptionId =
    process.env["DYNATRACE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DYNATRACE_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getVMHostPayload(resourceGroupName, monitorName);
  console.log(result);
}
