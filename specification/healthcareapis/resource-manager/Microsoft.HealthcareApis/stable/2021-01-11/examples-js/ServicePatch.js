const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function patchService() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const resourceName = "service1";
  const servicePatchDescription = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.beginUpdateAndWait(
    resourceGroupName,
    resourceName,
    servicePatchDescription
  );
  console.log(result);
}

patchService().catch(console.error);
