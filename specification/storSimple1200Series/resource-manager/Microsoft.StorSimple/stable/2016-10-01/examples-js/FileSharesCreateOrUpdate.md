```javascript
const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the file share.
 *
 * @summary Creates or updates the file share.
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/FileSharesCreateOrUpdate.json
 */
async function fileSharesCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-4XY4FI2IVG";
  const fileServerName = "HSDK-4XY4FI2IVG";
  const shareName = "Auto-TestFileShare1";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const fileShare = {
    name: "Auto-TestFileShare1",
    description: "Demo FileShare for SDK Test Tiered",
    adminUser: "fareastidcdlslb",
    dataPolicy: "Tiered",
    monitoringStatus: "Enabled",
    provisionedCapacityInBytes: 536870912000,
    shareStatus: "Online",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.fileShares.beginCreateOrUpdateAndWait(
    deviceName,
    fileServerName,
    shareName,
    resourceGroupName,
    managerName,
    fileShare
  );
  console.log(result);
}

fileSharesCreateOrUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple1200series_2.0.1/sdk/storsimple1200series/arm-storsimple1200series/README.md) on how to add the SDK to your project and authenticate.
