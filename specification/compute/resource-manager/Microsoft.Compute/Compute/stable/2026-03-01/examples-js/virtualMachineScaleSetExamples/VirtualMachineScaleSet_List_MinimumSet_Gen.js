const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of all VM scale sets under a resource group.
 *
 * @summary gets a list of all VM scale sets under a resource group.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_List_MinimumSet_Gen.json
 */
async function virtualMachineScaleSetListMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualMachineScaleSets.list("rgcompute")) {
    resArray.push(item);
  }

  console.log(resArray);
}
