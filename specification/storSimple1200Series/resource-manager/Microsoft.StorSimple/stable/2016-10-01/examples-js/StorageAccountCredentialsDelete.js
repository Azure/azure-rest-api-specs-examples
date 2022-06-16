const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function storageAccountCredentialsDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const credentialName = "DummySacForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.storageAccountCredentials.beginDeleteAndWait(
    credentialName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

storageAccountCredentialsDelete().catch(console.error);
