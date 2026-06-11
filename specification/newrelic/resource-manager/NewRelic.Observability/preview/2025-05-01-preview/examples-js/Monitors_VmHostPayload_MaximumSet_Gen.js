const { NewRelicObservability } = require("@azure/arm-newrelicobservability");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to returns the payload that needs to be passed in the request body for installing the New Relic agent on a VM, providing the necessary configuration details
 *
 * @summary returns the payload that needs to be passed in the request body for installing the New Relic agent on a VM, providing the necessary configuration details
 * x-ms-original-file: 2025-05-01-preview/Monitors_VmHostPayload_MaximumSet_Gen.json
 */
async function monitorsVmHostPayloadMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NewRelicObservability(credential, subscriptionId);
  const result = await client.monitors.vmHostPayload("rgopenapi", "ipxmlcbonyxtolzejcjshkmlron");
  console.log(result);
}
