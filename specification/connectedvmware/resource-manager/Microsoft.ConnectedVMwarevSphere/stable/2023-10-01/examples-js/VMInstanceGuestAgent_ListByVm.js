const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the list of GuestAgent of the given vm.
 *
 * @summary Returns the list of GuestAgent of the given vm.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/VMInstanceGuestAgent_ListByVm.json
 */
async function guestAgentListByVM() {
  const resourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential);
  const resArray = new Array();
  for await (let item of client.vMInstanceGuestAgents.list(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
