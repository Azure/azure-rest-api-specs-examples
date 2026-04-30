const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to changes ServiceState property for a given service
 *
 * @summary changes ServiceState property for a given service
 * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_SetOrchestrationServiceState_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetSetOrchestrationServiceStateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.virtualMachineScaleSets.setOrchestrationServiceState(
    "rgcompute",
    "aaaaaaaaaaaaaaaaaaaaaaaa",
    { serviceName: "AutomaticRepairs", action: "Resume" },
  );
}
