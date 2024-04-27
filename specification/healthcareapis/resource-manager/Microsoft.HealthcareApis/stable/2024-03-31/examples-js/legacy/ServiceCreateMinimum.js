const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the metadata of a service instance.
 *
 * @summary Create or update the metadata of a service instance.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/ServiceCreateMinimum.json
 */
async function createOrUpdateAServiceWithMinimumParameters() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "rg1";
  const resourceName = "service2";
  const serviceDescription = {
    kind: "fhir-R4",
    location: "westus2",
    properties: {
      accessPolicies: [{ objectId: "c487e7d1-3210-41a3-8ccc-e9372b78da47" }],
    },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    serviceDescription,
  );
  console.log(result);
}
