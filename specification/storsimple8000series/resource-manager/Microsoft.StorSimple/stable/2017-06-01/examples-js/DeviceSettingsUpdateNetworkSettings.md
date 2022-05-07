Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple8000series_2.0.1/sdk/storsimple8000series/arm-storsimple8000series/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { StorSimple8000SeriesManagementClient } = require("@azure/arm-storsimple8000series");
const { DefaultAzureCredential } = require("@azure/identity");

async function deviceSettingsUpdateNetworkSettings() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "Device05ForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "ManagerForSDKTest1";
  const parameters = {
    dnsSettings: {
      primaryDnsServer: "10.171.65.60",
      secondaryDnsServers: ["8.8.8.8"],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimple8000SeriesManagementClient(credential, subscriptionId);
  const result = await client.deviceSettings.beginUpdateNetworkSettingsAndWait(
    deviceName,
    resourceGroupName,
    managerName,
    parameters
  );
  console.log(result);
}

deviceSettingsUpdateNetworkSettings().catch(console.error);
```
