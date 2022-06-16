const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to retrieve SAS URIs for a virtual machine's boot diagnostic logs.
 *
 * @summary The operation to retrieve SAS URIs for a virtual machine's boot diagnostic logs.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachine_RetrieveBootDiagnosticsData.json
 */
async function retrieveBootDiagnosticsDataOfAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ResourceGroup";
  const vmName = "VMName";
  const sasUriExpirationTimeInMinutes = 60;
  const options = {
    sasUriExpirationTimeInMinutes,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.retrieveBootDiagnosticsData(
    resourceGroupName,
    vmName,
    options
  );
  console.log(result);
}

retrieveBootDiagnosticsDataOfAVirtualMachine().catch(console.error);
