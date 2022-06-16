const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Exposes the auto provisioning settings of the subscriptions
 *
 * @summary Exposes the auto provisioning settings of the subscriptions
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/AutoProvisioningSettings/GetAutoProvisioningSettingsSubscription_example.json
 */
async function getAutoProvisioningSettingsForSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.autoProvisioningSettings.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAutoProvisioningSettingsForSubscription().catch(console.error);
