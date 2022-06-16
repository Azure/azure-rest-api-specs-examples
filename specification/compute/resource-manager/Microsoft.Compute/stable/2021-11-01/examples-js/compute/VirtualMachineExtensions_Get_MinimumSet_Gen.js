const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineExtensionsGetMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "a";
  const vmExtensionName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.get(
    resourceGroupName,
    vmName,
    vmExtensionName
  );
  console.log(result);
}

virtualMachineExtensionsGetMinimumSetGen().catch(console.error);
