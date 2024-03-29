const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details of a specific setting
 *
 * @summary Details of a specific setting
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/AutoProvisioningSettings/GetAutoProvisioningSettingSubscription_example.json
 */
async function getAnAutoProvisioningSettingForSubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const settingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.autoProvisioningSettings.get(settingName);
  console.log(result);
}
