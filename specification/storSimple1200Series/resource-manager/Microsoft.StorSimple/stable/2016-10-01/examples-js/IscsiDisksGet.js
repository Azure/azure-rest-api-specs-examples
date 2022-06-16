const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function iscsiDisksGet() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-0NZI14MDTF";
  const iscsiServerName = "HSDK-0NZI14MDTF";
  const diskName = "Auto-TestIscsiDisk1";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.iscsiDisks.get(
    deviceName,
    iscsiServerName,
    diskName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

iscsiDisksGet().catch(console.error);
