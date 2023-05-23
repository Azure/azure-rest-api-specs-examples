const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Detach volume from the provided virtual machine.
 *
 * @summary Detach volume from the provided virtual machine.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/VirtualMachines_DetachVolume.json
 */
async function detachVolumeFromVirtualMachine() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const virtualMachineName = "virtualMachineName";
  const virtualMachineDetachVolumeParameters = {
    volumeId:
      "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.virtualMachines.beginDetachVolumeAndWait(
    resourceGroupName,
    virtualMachineName,
    virtualMachineDetachVolumeParameters
  );
  console.log(result);
}
