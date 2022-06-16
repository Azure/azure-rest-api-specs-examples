const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete the extension.
 *
 * @summary The operation to delete the extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineExtensions_Delete_MaximumSet_Gen.json
 */
async function virtualMachineExtensionsDeleteMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmName = "aaaaaaaaaaaaa";
  const vmExtensionName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.beginDeleteAndWait(
    resourceGroupName,
    vmName,
    vmExtensionName
  );
  console.log(result);
}

virtualMachineExtensionsDeleteMaximumSetGen().catch(console.error);
