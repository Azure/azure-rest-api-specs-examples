const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to assess patches on a vSphere VMware machine identity in Azure.
 *
 * @summary The operation to assess patches on a vSphere VMware machine identity in Azure.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/VirtualMachineAssessPatches.json
 */
async function assessPatchStateOfAMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroupName";
  const name = "myMachineName";
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.virtualMachines.beginAssessPatchesAndWait(resourceGroupName, name);
  console.log(result);
}

assessPatchStateOfAMachine().catch(console.error);
