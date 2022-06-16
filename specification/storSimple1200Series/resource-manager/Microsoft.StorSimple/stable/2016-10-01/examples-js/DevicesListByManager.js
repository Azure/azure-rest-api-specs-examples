const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all the devices in a manager.
 *
 * @summary Retrieves all the devices in a manager.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesListByManager.json
 */
async function devicesListByManager() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.devices.listByManager(resourceGroupName, managerName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

devicesListByManager().catch(console.error);
