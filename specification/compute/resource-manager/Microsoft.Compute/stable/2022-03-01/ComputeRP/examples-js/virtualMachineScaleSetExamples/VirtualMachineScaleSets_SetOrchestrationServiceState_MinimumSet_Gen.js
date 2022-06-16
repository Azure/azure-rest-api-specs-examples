const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Changes ServiceState property for a given service
 *
 * @summary Changes ServiceState property for a given service
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_SetOrchestrationServiceState_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetsSetOrchestrationServiceStateMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const parameters = {
    action: "Resume",
    serviceName: "AutomaticRepairs",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginSetOrchestrationServiceStateAndWait(
    resourceGroupName,
    vmScaleSetName,
    parameters
  );
  console.log(result);
}

virtualMachineScaleSetsSetOrchestrationServiceStateMinimumSetGen().catch(console.error);
