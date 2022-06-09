```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a security assessment on your resource. An assessment metadata that describes this assessment must be predefined with the same name before inserting the assessment result
 *
 * @summary Create a security assessment on your resource. An assessment metadata that describes this assessment must be predefined with the same name before inserting the assessment result
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/PutAssessment_example.json
 */
async function createSecurityRecommendationTaskOnAResource() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2";
  const assessmentName = "8bb8be0a-6010-4789-812f-e4d661c4ed0e";
  const assessment = {
    resourceDetails: { source: "Azure" },
    status: { code: "Healthy" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.assessments.createOrUpdate(resourceId, assessmentName, assessment);
  console.log(result);
}

createSecurityRecommendationTaskOnAResource().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
