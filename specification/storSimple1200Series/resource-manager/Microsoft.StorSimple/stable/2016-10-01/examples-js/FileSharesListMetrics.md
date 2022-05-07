Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple1200series_2.0.1/sdk/storsimple1200series/arm-storsimple1200series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the file share metrics
 *
 * @summary Gets the file share metrics
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesListMetrics.json
 */
async function fileSharesListMetrics() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-DMNJB2PET0";
  const fileServerName = "HSDK-DMNJB2PET0";
  const shareName = "Auto-TestFileShare2";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const filter = "startTime ge '2018-08-10T18:30:00Z' and endTime le '2018-08-11T18:30:00Z'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fileShares.listMetrics(
    deviceName,
    fileServerName,
    shareName,
    resourceGroupName,
    managerName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

fileSharesListMetrics().catch(console.error);
```
