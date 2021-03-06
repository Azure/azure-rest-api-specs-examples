const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function managersListMetrics() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const filter =
    "name/value eq 'PrimaryStorageTieredUsed' and timeGrain eq 'PT1H' and startTime ge '2017-06-17T18:30:00Z' and endTime le '2017-06-21T18:30:00Z' and category eq 'CapacityUtilization'";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managers.listMetrics(resourceGroupName, managerName, filter)) {
    resArray.push(item);
  }
  console.log(resArray);
}

managersListMetrics().catch(console.error);
