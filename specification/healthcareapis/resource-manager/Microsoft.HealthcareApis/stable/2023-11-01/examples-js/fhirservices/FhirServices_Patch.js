const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch FHIR Service details.
 *
 * @summary Patch FHIR Service details.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/fhirservices/FhirServices_Patch.json
 */
async function updateAFhirService() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "testRG";
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
    fhirservicePatchResource,
  );
  console.log(result);
}
