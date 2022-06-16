const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function jobsListByManager() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const filter = "jobType eq 'FailoverVolumeContainers'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.listByManager(resourceGroupName, managerName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

jobsListByManager().catch(console.error);
