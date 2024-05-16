const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a security assessment on your scanned resource
 *
 * @summary Get a security assessment on your scanned resource
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/GetAssessment_example.json
 */
async function getSecurityRecommendationTaskFromSecurityDataLocation() {
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/myRg/providers/Microsoft.Compute/virtualMachineScaleSets/vmss2";
  const assessmentName = "21300918-b2e3-0346-785f-c77ff57d243b";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.assessments.get(resourceId, assessmentName);
  console.log(result);
}
