const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the device metrics.
 *
 * @summary Retrieves the device metrics.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesListMetrics.json
 */
async function devicesListMetrics() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-UGU4PITWNI";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const filter = "startTime ge '2018-08-10T18:30:00Z' and endTime le '2018-08-11T18:30:00Z'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.devices.listMetrics(
    deviceName,
    resourceGroupName,
    managerName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

devicesListMetrics().catch(console.error);
