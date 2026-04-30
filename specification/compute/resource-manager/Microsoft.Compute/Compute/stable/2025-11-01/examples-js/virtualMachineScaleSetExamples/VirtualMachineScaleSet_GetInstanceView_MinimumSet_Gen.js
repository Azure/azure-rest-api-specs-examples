const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the status of a VM scale set instance.
 *
 * @summary gets the status of a VM scale set instance.
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetInstanceView_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetGetInstanceViewMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.getInstanceView(
    "rgcompute",
    "aaaaaaaaaaaaaaa",
  );
  console.log(result);
}
