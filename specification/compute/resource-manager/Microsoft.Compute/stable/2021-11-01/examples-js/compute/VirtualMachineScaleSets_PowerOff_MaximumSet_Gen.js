const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetsPowerOffMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaa";
  const skipShutdown = true;
  const vmInstanceIDs = {
    instanceIds: ["aaaaaaaaaaaaaaaaa"],
  };
  const options = { skipShutdown: skipShutdown, vmInstanceIDs: vmInstanceIDs };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginPowerOffAndWait(
    resourceGroupName,
    vmScaleSetName,
    options
  );
  console.log(result);
}

virtualMachineScaleSetsPowerOffMaximumSetGen().catch(console.error);
