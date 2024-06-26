const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to power off (stop) a virtual machine instance.
 *
 * @summary The operation to power off (stop) a virtual machine instance.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/StopVirtualMachineInstance.json
 */
async function stopVirtualMachine() {
  const resourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.HybridCompute/machines/DemoVM";
  const body = { skipShutdown: true };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential);
  const result = await client.virtualMachineInstances.beginStopAndWait(resourceUri, options);
  console.log(result);
}
