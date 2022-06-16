const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a private link resource that need to be created for a service.
 *
 * @summary Gets a private link resource that need to be created for a service.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/PrivateLinkResourceGet.json
 */
async function privateLinkResourcesGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgname";
  const resourceName = "service1";
  const groupName = "fhir";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, resourceName, groupName);
  console.log(result);
}

privateLinkResourcesGet().catch(console.error);
