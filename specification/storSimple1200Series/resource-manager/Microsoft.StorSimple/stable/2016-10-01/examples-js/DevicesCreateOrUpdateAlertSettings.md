```javascript
const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the alert settings
 *
 * @summary Creates or updates the alert settings
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/DevicesCreateOrUpdateAlertSettings.json
 */
async function devicesCreateOrUpdateAlertSettings() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const deviceName = "HSDK-T4ZA3EAJFR";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const alertSettings = {
    name: "default",
    additionalRecipientEmailList: ["testuser@abc.com"],
    alertNotificationCulture: "en-US",
    emailNotification: "Enabled",
    notificationToServiceOwners: "Disabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.devices.beginCreateOrUpdateAlertSettingsAndWait(
    deviceName,
    resourceGroupName,
    managerName,
    alertSettings
  );
  console.log(result);
}

devicesCreateOrUpdateAlertSettings().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storsimple1200series_2.0.1/sdk/storsimple1200series/arm-storsimple1200series/README.md) on how to add the SDK to your project and authenticate.
