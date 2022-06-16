const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function chapSettingsDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-WSJQERQW3F";
  const chapUserName = "ChapSettingForSDKForDelete";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.chapSettingsOperations.beginDeleteAndWait(
    deviceName,
    chapUserName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

chapSettingsDelete().catch(console.error);
