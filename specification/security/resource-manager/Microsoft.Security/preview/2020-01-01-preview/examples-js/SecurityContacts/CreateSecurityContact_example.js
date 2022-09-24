const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create security contact configurations for the subscription
 *
 * @summary Create security contact configurations for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/SecurityContacts/CreateSecurityContact_example.json
 */
async function createSecurityContactData() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const securityContactName = "default";
  const securityContact = {
    alertNotifications: { minimalSeverity: "Low", state: "On" },
    emails: "john@contoso.com;jane@contoso.com",
    notificationsByRole: { roles: ["Owner"], state: "On" },
    phone: "+214-2754038",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityContacts.create(securityContactName, securityContact);
  console.log(result);
}

createSecurityContactData().catch(console.error);
