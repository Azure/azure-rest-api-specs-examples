const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function jobsListByManager() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const filter = "jobType eq 'Backup'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.listByManager(resourceGroupName, managerName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

jobsListByManager().catch(console.error);
