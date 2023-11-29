const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the attached data networks associated with a packet core data plane.
 *
 * @summary Gets all the attached data networks associated with a packet core data plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/AttachedDataNetworkListByPacketCoreDataPlane.json
 */
async function listAttachedDataNetworksInADataPlane() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const packetCoreDataPlaneName = "TestPacketCoreDP";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.attachedDataNetworks.listByPacketCoreDataPlane(
    resourceGroupName,
    packetCoreControlPlaneName,
    packetCoreDataPlaneName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
