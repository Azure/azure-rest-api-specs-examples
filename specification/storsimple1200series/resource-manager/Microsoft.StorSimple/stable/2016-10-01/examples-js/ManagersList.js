const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all the managers in a subscription.
 *
 * @summary Retrieves all the managers in a subscription.
 * x-ms-original-file: specification/storsimple1200series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ManagersList.json
 */
async function managersList() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.managers.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

managersList().catch(console.error);
