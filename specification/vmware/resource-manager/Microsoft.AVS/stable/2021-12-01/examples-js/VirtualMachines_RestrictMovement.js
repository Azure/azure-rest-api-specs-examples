const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enable or disable DRS-driven VM movement restriction
 *
 * @summary Enable or disable DRS-driven VM movement restriction
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/VirtualMachines_RestrictMovement.json
 */
async function virtualMachineRestrictMovement() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const clusterName = "cluster1";
  const virtualMachineId = "vm-209";
  const restrictMovement = {
    restrictMovement: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.virtualMachines.beginRestrictMovementAndWait(
    resourceGroupName,
    privateCloudName,
    clusterName,
    virtualMachineId,
    restrictMovement
  );
  console.log(result);
}

virtualMachineRestrictMovement().catch(console.error);
