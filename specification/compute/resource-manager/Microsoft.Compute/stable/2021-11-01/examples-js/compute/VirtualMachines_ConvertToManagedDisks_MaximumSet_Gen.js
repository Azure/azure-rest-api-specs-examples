const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachinesConvertToManagedDisksMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginConvertToManagedDisksAndWait(
    resourceGroupName,
    vmName
  );
  console.log(result);
}

virtualMachinesConvertToManagedDisksMaximumSetGen().catch(console.error);
