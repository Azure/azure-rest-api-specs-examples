const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function iscsiDisksListMetricDefinition() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-WSJQERQW3F";
  const iscsiServerName = "HSDK-WSJQERQW3F";
  const diskName = "TieredIscsiDiskForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iscsiDisks.listMetricDefinition(
    deviceName,
    iscsiServerName,
    diskName,
    resourceGroupName,
    managerName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

iscsiDisksListMetricDefinition().catch(console.error);
