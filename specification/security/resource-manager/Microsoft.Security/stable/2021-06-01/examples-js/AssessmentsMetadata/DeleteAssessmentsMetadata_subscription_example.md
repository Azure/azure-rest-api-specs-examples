Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete metadata information on an assessment type in a specific subscription, will cause the deletion of all the assessments of that type in that subscription
 *
 * @summary Delete metadata information on an assessment type in a specific subscription, will cause the deletion of all the assessments of that type in that subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/DeleteAssessmentsMetadata_subscription_example.json
 */
async function deleteASecurityAssessmentMetadataForSubscription() {
  const subscriptionId = "0980887d-03d6-408c-9566-532f3456804e";
  const assessmentMetadataName = "ca039e75-a276-4175-aebc-bcd41e4b14b7";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.assessmentsMetadata.deleteInSubscription(assessmentMetadataName);
  console.log(result);
}

deleteASecurityAssessmentMetadataForSubscription().catch(console.error);
```
