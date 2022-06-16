const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function alertsListByManager() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const filter =
    "status eq 'Active' and appearedOnTime ge '2017-06-09T18:30:00Z' and appearedOnTime le '2017-06-19T18:30:00Z' and sourceType eq 'Device' and sourceName eq 'Device05ForSDKTest'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.alerts.listByManager(resourceGroupName, managerName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

alertsListByManager().catch(console.error);
