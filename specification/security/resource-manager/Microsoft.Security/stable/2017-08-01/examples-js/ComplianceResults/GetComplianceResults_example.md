Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Security Compliance Result
 *
 * @summary Security Compliance Result
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2017-08-01/examples/ComplianceResults/GetComplianceResults_example.json
 */
async function getComplianceResultsOnSubscription() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const complianceResultName = "DesignateMoreThanOneOwner";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.complianceResults.get(resourceId, complianceResultName);
  console.log(result);
}

getComplianceResultsOnSubscription().catch(console.error);
```
