const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a service.
 *
 * @summary Gets the private link resources that need to be created for a service.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/PrivateLinkResourcesListByService.json
 */
async function privateLinkResourcesListGroupIds() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgname";
  const resourceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.listByService(resourceGroupName, resourceName);
  console.log(result);
}

privateLinkResourcesListGroupIds().catch(console.error);
