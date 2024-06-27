const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the virtual machine instances within the specified parent resource.
 *
 * @summary Lists all of the virtual machine instances within the specified parent resource.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/VirtualMachineInstances_List_MinimumSet_Gen.json
 */
async function virtualMachineInstancesListMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const resArray = new Array();
  for await (let item of client.virtualMachineInstances.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
