Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Supported regulatory compliance standards details and state
 *
 * @summary Supported regulatory compliance standards details and state
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceStandardList_example.json
 */
async function getAllSupportedRegulatoryComplianceStandardsDetailsAndState() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.regulatoryComplianceStandards.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllSupportedRegulatoryComplianceStandardsDetailsAndState().catch(console.error);
```
