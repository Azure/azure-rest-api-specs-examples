Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Security contact configurations for the subscription
 *
 * @summary Security contact configurations for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/SecurityContacts/UpdateSecurityContact_full_example.json
 */
async function updateSecurityContactDataFull() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const securityContactName = "john";
  const securityContact = {
    name: "default1",
    type: "Microsoft.Security/securityContacts",
    alertNotifications: "On",
    id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/securityContacts/default1",
    phone: "(214)275-4038",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityContacts.update(securityContactName, securityContact);
  console.log(result);
}

updateSecurityContactDataFull().catch(console.error);
```
