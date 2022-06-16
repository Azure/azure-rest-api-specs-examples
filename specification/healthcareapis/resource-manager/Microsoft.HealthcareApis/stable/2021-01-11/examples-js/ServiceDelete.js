const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const resourceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.beginDeleteAndWait(resourceGroupName, resourceName);
  console.log(result);
}

deleteService().catch(console.error);
