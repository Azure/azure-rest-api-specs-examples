const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function alertsListByManager() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hManagerForSDKTest4";
  const filter = "status eq 'Active' and severity eq 'Critical'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.alerts.listByManager(resourceGroupName, managerName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

alertsListByManager().catch(console.error);
