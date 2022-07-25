const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the data networks associated with a packet core data plane.
 *
 * @summary Gets all the data networks associated with a packet core data plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/AttachedDataNetworkListByPacketCoreDataPlane.json
 */
async function listAttachedDataNetworksInADataPlane() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
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

listAttachedDataNetworksInADataPlane().catch(console.error);
