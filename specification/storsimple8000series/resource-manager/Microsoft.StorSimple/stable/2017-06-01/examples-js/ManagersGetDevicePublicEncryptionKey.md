```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the public encryption key of the device.
 *
 * @summary Returns the public encryption key of the device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/ManagersGetDevicePublicEncryptionKey.json
 */
async function managersGetDevicePublicEncryptionKey() {
  const subscriptionId = "d3ebfe71-b7a9-4c57-92b9-68a2afde4de5";
  const deviceName = "sca01forsdktest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.managers.getDevicePublicEncryptionKey(
    deviceName,
    resourceGroupName,
    managerName
  );
  console.log(result);
}

managersGetDevicePublicEncryptionKey().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.
