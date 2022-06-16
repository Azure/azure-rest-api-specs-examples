const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Backup the iSCSI server now.
 *
 * @summary Backup the iSCSI server now.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/IscsiServersBackupNow.json
 */
async function iscsiServersBackupNow() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-0NZI14MDTF";
  const iscsiServerName = "HSDK-0NZI14MDTF";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.iscsiServers.beginBackupNowAndWait(
    deviceName,
    iscsiServerName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

iscsiServersBackupNow().catch(console.error);
