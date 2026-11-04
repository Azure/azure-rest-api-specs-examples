const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates an existing AutoscaleSettingsResource. To update other fields use the CreateOrUpdate method.
 *
 * @summary updates an existing AutoscaleSettingsResource. To update other fields use the CreateOrUpdate method.
 * x-ms-original-file: 2025-09-01-preview/patchAutoscaleSettingRemoveIdentity.json
 */
async function removeTheManagedIdentityFromAnAutoscaleSetting() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.autoscaleSettings.update("TestingMetricsScaleSet", "MySetting", {
    identity: { type: "None" },
  });
  console.log(result);
}
