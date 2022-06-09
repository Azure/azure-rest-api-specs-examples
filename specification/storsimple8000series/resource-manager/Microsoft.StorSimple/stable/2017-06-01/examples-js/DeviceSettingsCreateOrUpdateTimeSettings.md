```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the time settings of the specified device.
 *
 * @summary Creates or updates the time settings of the specified device.
 * x-ms-original-file: specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsCreateOrUpdateTimeSettings.json
 */
async function deviceSettingsCreateOrUpdateTimeSettings() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    primaryTimeServer: "time.windows.com",
    secondaryTimeServer: ["8.8.8.8"],
    timeZone: "Pacific Standard Time",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.deviceSettings.beginCreateOrUpdateTimeSettingsAndWait(
    deviceName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

deviceSettingsCreateOrUpdateTimeSettings().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.
