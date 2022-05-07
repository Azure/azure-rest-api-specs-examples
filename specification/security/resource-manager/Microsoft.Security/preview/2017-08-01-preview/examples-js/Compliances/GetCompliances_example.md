Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Compliance scores of the specific management group.
 *
 * @summary The Compliance scores of the specific management group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliances_example.json
 */
async function getSecurityComplianceDataOverTime() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.compliances.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getSecurityComplianceDataOverTime().catch(console.error);
```
