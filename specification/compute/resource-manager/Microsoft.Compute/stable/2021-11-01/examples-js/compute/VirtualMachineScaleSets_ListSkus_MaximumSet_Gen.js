const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of SKUs available for your VM scale set, including the minimum and maximum VM instances allowed for each SKU.
 *
 * @summary Gets a list of SKUs available for your VM scale set, including the minimum and maximum VM instances allowed for each SKU.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_ListSkus_MaximumSet_Gen.json
 */
async function virtualMachineScaleSetsListSkusMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineScaleSets.listSkus(
    resourceGroupName,
    vmScaleSetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineScaleSetsListSkusMaximumSetGen().catch(console.error);
