const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all virtual machines in a VM scale sets.
 *
 * @summary Gets a list of all virtual machines in a VM scale sets.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetVMListMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const virtualMachineScaleSetName = "aaaaaaaaaaaaaaaaaaaaaa";
  const filter = "aaaaaaaaaaaaaa";
  const select = "aaaaaaaaaaaaaaaaaaaaa";
  const expand = "aaaaaaaaaaaaa";
  const options = {
    filter,
    select,
    expand,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineScaleSetVMs.list(
    resourceGroupName,
    virtualMachineScaleSetName,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
