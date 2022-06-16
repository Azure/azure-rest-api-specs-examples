const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineExtensionsGetMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const vmExtensionName = "aaaaaaa";
  const expand = "aaaaaa";
  const options = { expand: expand };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.get(
    resourceGroupName,
    vmName,
    vmExtensionName,
    options
  );
  console.log(result);
}

virtualMachineExtensionsGetMaximumSetGen().catch(console.error);
