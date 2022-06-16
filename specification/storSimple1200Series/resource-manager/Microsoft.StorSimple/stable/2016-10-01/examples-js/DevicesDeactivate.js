const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function devicesDeactivate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "SDK-DELETE";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForDeleteOperation";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.devices.beginDeactivateAndWait(
    deviceName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

devicesDeactivate().catch(console.error);
