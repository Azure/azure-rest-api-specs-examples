const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function fileSharesGet() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-4XY4FI2IVG";
  const fileServerName = "HSDK-4XY4FI2IVG";
  const shareName = "Auto-TestFileShare1";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.fileShares.get(
    deviceName,
    fileServerName,
    shareName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

fileSharesGet().catch(console.error);
