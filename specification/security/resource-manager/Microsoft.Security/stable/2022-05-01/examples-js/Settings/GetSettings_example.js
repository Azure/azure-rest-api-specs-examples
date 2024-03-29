const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Settings about different configurations in Microsoft Defender for Cloud
 *
 * @summary Settings about different configurations in Microsoft Defender for Cloud
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-05-01/examples/Settings/GetSettings_example.json
 */
async function getSettingsOfSubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.settings.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
