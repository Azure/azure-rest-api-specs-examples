const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Import a virtual machine into a different lab. This operation can take a while to complete.
 *
 * @summary Import a virtual machine into a different lab. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ImportVirtualMachine.json
 */
async function labsImportVirtualMachine() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const importLabVirtualMachineRequest = {
    destinationVirtualMachineName: "{vmName}",
    sourceVirtualMachineResourceId:
      "/subscriptions/{subscriptionId}/resourceGroups/{otherResourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.labs.beginImportVirtualMachineAndWait(
    resourceGroupName,
    name,
    importLabVirtualMachineRequest
  );
  console.log(result);
}

labsImportVirtualMachine().catch(console.error);
