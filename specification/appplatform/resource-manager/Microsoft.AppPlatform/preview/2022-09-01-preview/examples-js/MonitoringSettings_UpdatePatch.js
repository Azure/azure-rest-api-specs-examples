const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the Monitoring Setting.
 *
 * @summary Update the Monitoring Setting.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/MonitoringSettings_UpdatePatch.json
 */
async function monitoringSettingsUpdatePatch() {
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
  const result = await client.monitoringSettings.beginUpdatePatchAndWait(
    resourceGroupName,
    serviceName,
    monitoringSettingResource
  );
  console.log(result);
}

monitoringSettingsUpdatePatch().catch(console.error);
