const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetsGetInstanceViewMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSets.getInstanceView(
    resourceGroupName,
    vmScaleSetName
  );
  console.log(result);
}

virtualMachineScaleSetsGetInstanceViewMaximumSetGen().catch(console.error);
