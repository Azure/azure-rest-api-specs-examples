const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function iscsiServersDelete() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-UGU4PITWNI";
  const iscsiServerName = "HSDK-UGU4PITWNI";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.iscsiServers.beginDeleteAndWait(
    deviceName,
    iscsiServerName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

iscsiServersDelete().catch(console.error);
