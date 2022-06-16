const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Security contact configurations for the subscription
 *
 * @summary Security contact configurations for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/SecurityContacts/UpdateSecurityContact_min_example.json
 */
async function updateSecurityContactDataMinimal() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const securityContactName = "default2";
  const securityContact = {
    name: "default2",
    type: "Microsoft.Security/securityContacts",
    alertNotifications: "On",
    id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/securityContacts/default2",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityContacts.update(securityContactName, securityContact);
  console.log(result);
}

updateSecurityContactDataMinimal().catch(console.error);
