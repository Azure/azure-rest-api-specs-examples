Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get metadata information on an assessment type in a specific subscription
 *
 * @summary Get metadata information on an assessment type in a specific subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_subscription_example.json
 */
async function getSecurityAssessmentMetadataForSubscription() {
  const subscriptionId = "0980887d-03d6-408c-9566-532f3456804e";
  const assessmentMetadataName = "21300918-b2e3-0346-785f-c77ff57d243b";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.assessmentsMetadata.getInSubscription(assessmentMetadataName);
  console.log(result);
}

getSecurityAssessmentMetadataForSubscription().catch(console.error);
```
