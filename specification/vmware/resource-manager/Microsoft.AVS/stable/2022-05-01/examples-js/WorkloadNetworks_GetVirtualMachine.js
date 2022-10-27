const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a virtual machine by id in a private cloud workload network.
 *
 * @summary Get a virtual machine by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetVirtualMachine.json
 */
async function workloadNetworksGetVirtualMachine() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const virtualMachineId = "vm1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.getVirtualMachine(
    resourceGroupName,
    privateCloudName,
    virtualMachineId
  );
  console.log(result);
}

workloadNetworksGetVirtualMachine().catch(console.error);
