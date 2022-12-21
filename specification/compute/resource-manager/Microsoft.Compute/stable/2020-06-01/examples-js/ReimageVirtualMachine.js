const { ComputeManagementClient } = require("@azure/arm-compute-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reimages the virtual machine which has an ephemeral OS disk back to its initial state.
 *
 * @summary Reimages the virtual machine which has an ephemeral OS disk back to its initial state.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-06-01/examples/ReimageVirtualMachine.json
 */
async function reimageAVirtualMachine() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVMName";
  const parameters = { tempDisk: true };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginReimageAndWait(
    resourceGroupName,
    vmName,
    options
  );
  console.log(result);
}
