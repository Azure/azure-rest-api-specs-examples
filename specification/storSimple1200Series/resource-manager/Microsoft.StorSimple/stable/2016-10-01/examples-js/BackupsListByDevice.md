Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple1200series_2.0.1/sdk/storsimple1200series/arm-storsimple1200series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all the backups in a device. Can be used to get the backups for failover also.
 *
 * @summary Retrieves all the backups in a device. Can be used to get the backups for failover also.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/BackupsListByDevice.json
 */
async function backupsListByDevice() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-4XY4FI2IVG";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const filter = "initiatedBy eq 'Manual'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.backups.listByDevice(
    deviceName,
    resourceGroupName,
    managerName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

backupsListByDevice().catch(console.error);
```
