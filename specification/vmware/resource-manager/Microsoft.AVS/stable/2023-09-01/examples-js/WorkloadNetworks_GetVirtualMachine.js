const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a WorkloadNetworkVirtualMachine
 *
 * @summary Get a WorkloadNetworkVirtualMachine
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_GetVirtualMachine.json
 */
async function workloadNetworksGetVirtualMachine() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const virtualMachineId = "vm1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.getVirtualMachine(
    resourceGroupName,
    privateCloudName,
    virtualMachineId,
  );
  console.log(result);
}
