Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Selected regulatory compliance control details and state
 *
 * @summary Selected regulatory compliance control details and state
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/RegulatoryCompliance/getRegulatoryComplianceControl_example.json
 */
async function getSelectedRegulatoryComplianceControlDetailsAndState() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const regulatoryComplianceStandardName = "PCI-DSS-3.2";
  const regulatoryComplianceControlName = "1.1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.regulatoryComplianceControls.get(
    regulatoryComplianceStandardName,
    regulatoryComplianceControlName
  );
  console.log(result);
}

getSelectedRegulatoryComplianceControlDetailsAndState().catch(console.error);
```
