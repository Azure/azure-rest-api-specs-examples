const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

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
