const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function iscsiServersCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-WSJQERQW3F";
  const iscsiServerName = "HSDK-WSJQERQW3F";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const iscsiServer = {
    name: "HSDK-WSJQERQW3F",
    type: "Microsoft.StorSimple/managers/devices/iscsiServers",
    backupScheduleGroupId:
      "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/backupScheduleGroups/Default-HSDK-WSJQERQW3F-BackupScheduleGroup",
    chapId:
      "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/chapSettings/ChapSettingForSDK",
    id: "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/devices/HSDK-WSJQERQW3F/iscsiServers/HSDK-WSJQERQW3F",
    reverseChapId: "",
    storageDomainId:
      "/subscriptions/9eb689cd-7243-43b4-b6f6-5c65cb296641/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/hAzureSDKOperations/storageDomains/Default-HSDK-WSJQERQW3F-StorageDomain",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.iscsiServers.beginCreateOrUpdateAndWait(
    deviceName,
    iscsiServerName,
    resourceGroupName,
    managerName,
    iscsiServer
  );
  console.log(result);
}

iscsiServersCreateOrUpdate().catch(console.error);
