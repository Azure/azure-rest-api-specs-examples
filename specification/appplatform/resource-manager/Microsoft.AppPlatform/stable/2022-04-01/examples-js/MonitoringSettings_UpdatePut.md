```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the Monitoring Setting.
 *
 * @summary Update the Monitoring Setting.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/MonitoringSettings_UpdatePut.json
 */
async function monitoringSettingsUpdatePut() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const monitoringSettingResource = {
    properties: {
      appInsightsInstrumentationKey: "00000000-0000-0000-0000-000000000000",
      appInsightsSamplingRate: 10,
      traceEnabled: true,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.monitoringSettings.beginUpdatePutAndWait(
    resourceGroupName,
    serviceName,
    monitoringSettingResource
  );
  console.log(result);
}

monitoringSettingsUpdatePut().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.
