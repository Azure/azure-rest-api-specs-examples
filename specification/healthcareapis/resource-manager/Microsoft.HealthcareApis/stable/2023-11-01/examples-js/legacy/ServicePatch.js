const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the metadata of a service instance.
 *
 * @summary Update the metadata of a service instance.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/legacy/ServicePatch.json
 */
async function patchService() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "rg1";
  const resourceName = "service1";
  const servicePatchDescription = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.beginUpdateAndWait(
    resourceGroupName,
    resourceName,
    servicePatchDescription,
  );
  console.log(result);
}
