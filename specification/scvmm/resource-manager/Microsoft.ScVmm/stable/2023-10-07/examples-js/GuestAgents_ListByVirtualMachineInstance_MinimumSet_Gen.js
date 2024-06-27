const { ScVmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of GuestAgent of the given vm.
 *
 * @summary Returns the list of GuestAgent of the given vm.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/GuestAgents_ListByVirtualMachineInstance_MinimumSet_Gen.json
 */
async function guestAgentsListByVirtualMachineInstanceMinimumSet() {
  const resourceUri = "gtgclehcbsyave";
  const credential = new DefaultAzureCredential();
  const client = new ScVmm(credential);
  const resArray = new Array();
  for await (let item of client.guestAgents.listByVirtualMachineInstance(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
