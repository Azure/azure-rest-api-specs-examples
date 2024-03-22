const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of the subvolumes in the volume
 *
 * @summary Returns a list of the subvolumes in the volume
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/Subvolumes_List.json
 */
async function subvolumesList() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.subvolumes.listByVolume(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
