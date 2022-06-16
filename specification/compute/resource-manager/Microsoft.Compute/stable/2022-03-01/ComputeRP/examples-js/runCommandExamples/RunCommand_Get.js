const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets specific run command for a subscription in a location.
 *
 * @summary Gets specific run command for a subscription in a location.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/RunCommand_Get.json
 */
async function virtualMachineRunCommandGet() {
  const subscriptionId = "24fb23e3-6ba3-41f0-9b6e-e41131d5d61e";
  const location = "SoutheastAsia";
  const commandId = "RunPowerShellScript";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineRunCommands.get(location, commandId);
  console.log(result);
}

virtualMachineRunCommandGet().catch(console.error);
