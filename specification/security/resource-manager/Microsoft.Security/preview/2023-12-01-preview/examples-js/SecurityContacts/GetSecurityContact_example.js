const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Default Security contact configurations for the subscription
 *
 * @summary Get Default Security contact configurations for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-12-01-preview/examples/SecurityContacts/GetSecurityContact_example.json
 */
async function getASecurityContact() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const securityContactName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityContacts.get(securityContactName);
  console.log(result);
}
