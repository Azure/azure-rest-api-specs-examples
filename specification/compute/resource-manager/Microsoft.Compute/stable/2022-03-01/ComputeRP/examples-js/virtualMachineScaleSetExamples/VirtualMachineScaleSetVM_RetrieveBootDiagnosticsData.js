const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to retrieve SAS URIs of boot diagnostic logs for a virtual machine in a VM scale set.
 *
 * @summary The operation to retrieve SAS URIs of boot diagnostic logs for a virtual machine in a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_RetrieveBootDiagnosticsData.json
 */
async function retrieveBootDiagnosticsDataOfAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const sasUriExpirationTimeInMinutes = 60;
  const options = {
    sasUriExpirationTimeInMinutes,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.retrieveBootDiagnosticsData(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    options
  );
  console.log(result);
}

retrieveBootDiagnosticsDataOfAVirtualMachine().catch(console.error);
