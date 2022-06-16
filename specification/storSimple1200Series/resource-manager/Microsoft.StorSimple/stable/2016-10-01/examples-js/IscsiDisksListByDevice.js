const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all the iSCSI disks in a device.
 *
 * @summary Retrieves all the iSCSI disks in a device.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiDisksListByDevice.json
 */
async function iscsiDisksListByDevice() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-0NZI14MDTF";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iscsiDisks.listByDevice(
    deviceName,
    resourceGroupName,
    managerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

iscsiDisksListByDevice().catch(console.error);
