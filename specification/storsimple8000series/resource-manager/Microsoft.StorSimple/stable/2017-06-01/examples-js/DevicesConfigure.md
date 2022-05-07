Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function devicesConfigure() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    currentDeviceName: "Device001ForSDKTest",
    friendlyName: "Device001ForSDKTest",
    networkInterfaceData0Settings: {
      controllerOneIp: "10.168.220.228",
      controllerZeroIp: "10.168.220.227",
    },
    timeZone: "Pacific Standard Time",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.devices.beginConfigureAndWait(
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

devicesConfigure().catch(console.error);
```
