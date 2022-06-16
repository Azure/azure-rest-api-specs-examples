const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetsPowerOffMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "a";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.beginPowerOffAndWait(
    resourceGroupName,
    vmScaleSetName
  );
  console.log(result);
}

virtualMachineScaleSetsPowerOffMinimumSetGen().catch(console.error);
