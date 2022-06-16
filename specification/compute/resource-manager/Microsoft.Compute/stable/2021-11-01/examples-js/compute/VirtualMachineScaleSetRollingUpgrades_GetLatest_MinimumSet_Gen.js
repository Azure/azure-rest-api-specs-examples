const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetRollingUpgradesGetLatestMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetRollingUpgrades.getLatest(
    resourceGroupName,
    vmScaleSetName
  );
  console.log(result);
}

virtualMachineScaleSetRollingUpgradesGetLatestMinimumSetGen().catch(console.error);
