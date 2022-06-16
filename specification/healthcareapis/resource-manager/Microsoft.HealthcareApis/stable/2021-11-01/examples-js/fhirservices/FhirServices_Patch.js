const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch FHIR Service details.
 *
 * @summary Patch FHIR Service details.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Patch.json
 */
async function updateAFhirService() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const fhirServiceName = "fhirservice1";
  const workspaceName = "workspace1";
  const fhirservicePatchResource = {
    tags: { tagKey: "tagValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.fhirServices.beginUpdateAndWait(
    resourceGroupName,
    fhirServiceName,
    workspaceName,
    fhirservicePatchResource
  );
  console.log(result);
}

updateAFhirService().catch(console.error);
