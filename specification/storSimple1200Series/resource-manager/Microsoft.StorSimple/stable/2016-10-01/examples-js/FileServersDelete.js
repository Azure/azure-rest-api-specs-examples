const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function fileServersDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-DMNJB2PET0";
  const fileServerName = "HSDK-DMNJB2PET0";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.fileServers.beginDeleteAndWait(
    deviceName,
    fileServerName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

fileServersDelete().catch(console.error);
