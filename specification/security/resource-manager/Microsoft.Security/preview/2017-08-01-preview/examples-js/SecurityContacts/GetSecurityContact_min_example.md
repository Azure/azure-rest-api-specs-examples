Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Security contact configurations for the subscription
 *
 * @summary Security contact configurations for the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/SecurityContacts/GetSecurityContact_min_example.json
 */
async function getSecurityContactDataMinimal() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const securityContactName = "default2";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityContacts.get(securityContactName);
  console.log(result);
}

getSecurityContactDataMinimal().catch(console.error);
```
