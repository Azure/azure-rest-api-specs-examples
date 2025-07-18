const { ChaosManagementClient } = require("@azure/arm-chaos");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a Target resource that extends a tracked regional resource.
 *
 * @summary get a Target resource that extends a tracked regional resource.
 * x-ms-original-file: 2025-01-01/Targets_Get.json
 */
async function getATargetThatExtendsAVirtualMachineResource() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
  const client = new ChaosManagementClient(credential, subscriptionId);
  const result = await client.targets.get(
    "exampleRG",
    "Microsoft.Compute",
    "virtualMachines",
    "exampleVM",
    "Microsoft-Agent",
  );
  console.log(result);
}
