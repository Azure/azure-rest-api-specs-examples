const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetsSetOrchestrationServiceStateMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaa";
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

virtualMachineScaleSetsSetOrchestrationServiceStateMaximumSetGen().catch(console.error);
