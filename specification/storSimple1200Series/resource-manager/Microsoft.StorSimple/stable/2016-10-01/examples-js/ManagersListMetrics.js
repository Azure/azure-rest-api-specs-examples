const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

async function managersListMetrics() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const filter = "startTime ge '2018-08-04T18:30:00Z' and endTime le '2018-08-11T18:30:00Z'";
  const options = { filter: filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managers.listMetrics(resourceGroupName, managerName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

managersListMetrics().catch(console.error);
