const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a FHIR Service.
 *
 * @summary Deletes a FHIR Service.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Delete.json
 */
async function deleteAFhirService() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const fhirServiceName = "fhirservice1";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.fhirServices.beginDeleteAndWait(
    resourceGroupName,
    fhirServiceName,
    workspaceName
  );
  console.log(result);
}

deleteAFhirService().catch(console.error);
