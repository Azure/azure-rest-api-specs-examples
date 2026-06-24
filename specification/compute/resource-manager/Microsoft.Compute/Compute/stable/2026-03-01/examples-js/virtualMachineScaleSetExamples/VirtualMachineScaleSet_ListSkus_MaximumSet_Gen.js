const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a list of SKUs available for your VM scale set, including the minimum and maximum VM instances allowed for each SKU.
 *
 * @summary gets a list of SKUs available for your VM scale set, including the minimum and maximum VM instances allowed for each SKU.
 * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListSkus_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetListSkusMaximumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.virtualMachineScaleSets.listSkus("rgcompute", "aaaaaa")) {
    resArray.push(item);
  }

  console.log(resArray);
}
