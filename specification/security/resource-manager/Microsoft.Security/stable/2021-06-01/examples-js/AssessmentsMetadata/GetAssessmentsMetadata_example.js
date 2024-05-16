const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get metadata information on an assessment type
 *
 * @summary Get metadata information on an assessment type
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/GetAssessmentsMetadata_example.json
 */
async function getSecurityAssessmentMetadata() {
  const assessmentMetadataName = "21300918-b2e3-0346-785f-c77ff57d243b";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.assessmentsMetadata.get(assessmentMetadataName);
  console.log(result);
}
