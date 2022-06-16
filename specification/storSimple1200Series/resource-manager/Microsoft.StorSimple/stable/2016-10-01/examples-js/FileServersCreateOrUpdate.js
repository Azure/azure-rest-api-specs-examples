const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function fileServersCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-4XY4FI2IVG";
  const fileServerName = "HSDK-4XY4FI2IVG";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const fileServer = {
    name: "HSDK-4XY4FI2IVG",
    description: "Demo FileServer for SDK Test",
    backupScheduleGroupId:
      "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/hsdk-4xy4fi2ivg/backupScheduleGroups/BackupSchGroupForSDKTest",
    domainName: "fareast.corp.microsoft.com",
    storageDomainId:
      "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/sd-fs-HSDK-4XY4FI2IVG",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.fileServers.beginCreateOrUpdateAndWait(
    deviceName,
    fileServerName,
    resourceGroupName,
    managerName,
    fileServer
  );
  console.log(result);
}

fileServersCreateOrUpdate().catch(console.error);
