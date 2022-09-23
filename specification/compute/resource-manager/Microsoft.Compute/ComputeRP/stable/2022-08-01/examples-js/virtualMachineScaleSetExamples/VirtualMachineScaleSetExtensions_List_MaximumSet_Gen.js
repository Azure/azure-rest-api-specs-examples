const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all extensions in a VM scale set.
 *
 * @summary Gets a list of all extensions in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetExtensions_List_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetExtensionsListMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineScaleSetExtensions.list(
    resourceGroupName,
    vmScaleSetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineScaleSetExtensionsListMaximumSetGen().catch(console.error);
