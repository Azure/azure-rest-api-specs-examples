const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified attached data network.
 *
 * @summary Deletes the specified attached data network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/AttachedDataNetworkDelete.json
 */
async function deleteAttachedDataNetworkResource() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const packetCoreDataPlaneName = "TestPacketCoreDP";
  const attachedDataNetworkName = "TestAttachedDataNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.attachedDataNetworks.beginDeleteAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName,
    attachedDataNetworkName,
  );
  console.log(result);
}
