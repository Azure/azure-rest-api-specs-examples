const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list with a single sensitivity settings resource
 *
 * @summary Gets a list with a single sensitivity settings resource
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-02-15-preview/examples/SensitivitySettings/GetSensitivitySettingsList_example.json
 */
async function getSensitivitySettingsList() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.sensitivitySettings.list();
  console.log(result);
}
