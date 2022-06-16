const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the metadata of a service instance.
 *
 * @summary Get the metadata of a service instance.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceGet.json
 */
async function getMetadata() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const resourceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.get(resourceGroupName, resourceName);
  console.log(result);
}

getMetadata().catch(console.error);
