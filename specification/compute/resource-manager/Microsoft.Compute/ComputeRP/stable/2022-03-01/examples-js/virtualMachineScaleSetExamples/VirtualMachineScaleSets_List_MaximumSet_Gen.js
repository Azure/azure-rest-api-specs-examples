const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all VM scale sets under a resource group.
 *
 * @summary Gets a list of all VM scale sets under a resource group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_List_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetsListMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineScaleSets.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineScaleSetsListMaximumSetGen().catch(console.error);
