const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates data sensitivity settings for sensitive data discovery
 *
 * @summary Updates data sensitivity settings for sensitive data discovery
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/SensitivitySettings/PutSensitivitySettings_example.json
 */
async function updateSensitivitySettings() {
  const sensitivitySettings = {
    sensitiveInfoTypesIds: [
      "f2f8a7a1-28c0-404b-9ab4-30a0a7af18cb",
      "b452f22b-f87d-4f48-8490-ecf0873325b5",
      "d59ee8b6-2618-404b-a5e7-aa377cd67543",
    ],
    sensitivityThresholdLabelId: "f2f8a7a1-28c0-404b-9ab4-30a0a7af18cb",
    sensitivityThresholdLabelOrder: 2,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.updateSensitivitySettings(sensitivitySettings);
  console.log(result);
}
