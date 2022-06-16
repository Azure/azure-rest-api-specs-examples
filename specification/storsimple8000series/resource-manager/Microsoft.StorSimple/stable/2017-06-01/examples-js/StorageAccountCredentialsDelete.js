const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCredentialsDelete() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const storageAccountCredentialName = "SACForTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.storageAccountCredentials.beginDeleteAndWait(
    storageAccountCredentialName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

storageAccountCredentialsDelete().catch(console.error);
