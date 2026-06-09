const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update the metadata of a service instance.
 *
 * @summary create or update the metadata of a service instance.
 * x-ms-original-file: 2025-04-01-preview/legacy/ServiceCreateMinimum.json
 */
async function createOrUpdateAServiceWithMinimumParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.createOrUpdate("rg1", "service2", {
    kind: "fhir-R4",
    location: "westus2",
    properties: { accessPolicies: [{ objectId: "c487e7d1-3210-41a3-8ccc-e9372b78da47" }] },
    tags: {},
  });
  console.log(result);
}
