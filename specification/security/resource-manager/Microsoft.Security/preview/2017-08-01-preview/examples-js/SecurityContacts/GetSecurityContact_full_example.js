const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Security contact configurations for the subscription
 *
 * @summary Security contact configurations for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/SecurityContacts/GetSecurityContact_full_example.json
 */
async function getSecurityContactDataFull() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const securityContactName = "default1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityContacts.get(securityContactName);
  console.log(result);
}

getSecurityContactDataFull().catch(console.error);
