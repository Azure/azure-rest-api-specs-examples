const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get metadata information on all assessment types
 *
 * @summary Get metadata information on all assessment types
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/ListAssessmentsMetadata_example.json
 */
async function listSecurityAssessmentMetadata() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (let item of client.assessmentsMetadata.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
