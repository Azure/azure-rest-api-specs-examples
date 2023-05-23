const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the provided rack.
All customer initiated requests will be rejected as the life cycle of this resource is managed by the system.
 *
 * @summary Delete the provided rack.
All customer initiated requests will be rejected as the life cycle of this resource is managed by the system.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/Racks_Delete.json
 */
async function deleteRack() {
  const subscriptionId = process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const rackName = "rackName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.racks.beginDeleteAndWait(resourceGroupName, rackName);
  console.log(result);
}
