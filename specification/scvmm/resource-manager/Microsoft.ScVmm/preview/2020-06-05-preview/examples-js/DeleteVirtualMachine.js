const { Scvmm } = require("@azure/arm-scvmm");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a VirtualMachine deployed on ScVmm fabric.
 *
 * @summary Deletes a VirtualMachine deployed on ScVmm fabric.
 * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVirtualMachine.json
 */
async function deleteVirtualMachine() {
  const subscriptionId =
    process.env["SCVMM_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["SCVMM_RESOURCE_GROUP"] || "testrg";
  const virtualMachineName = "DemoVM";
  const credential = new DefaultAzureCredential();
  const client = new Scvmm(credential, subscriptionId);
  const result = await client.virtualMachines.beginDeleteAndWait(
    resourceGroupName,
    virtualMachineName
  );
  console.log(result);
}
