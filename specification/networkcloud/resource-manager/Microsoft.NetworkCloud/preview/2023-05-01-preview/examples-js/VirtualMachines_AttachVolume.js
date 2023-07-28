const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Attach volume to the provided virtual machine.
 *
 * @summary Attach volume to the provided virtual machine.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/VirtualMachines_AttachVolume.json
 */
async function attachVolumeToVirtualMachine() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const virtualMachineName = "virtualMachineName";
  const virtualMachineAttachVolumeParameters = {
    volumeId:
      "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.virtualMachines.beginAttachVolumeAndWait(
    resourceGroupName,
    virtualMachineName,
    virtualMachineAttachVolumeParameters
  );
  console.log(result);
}
