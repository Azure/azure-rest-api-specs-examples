const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function backupsDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-UGU4PITWNI";
  const backupName = "315c3461-96ad-41bf-af0b-3a8be5113203";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.backups.beginDeleteAndWait(
    deviceName,
    backupName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

backupsDelete().catch(console.error);
