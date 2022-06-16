const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function checkNameAvailability() {
  const subscriptionId = "subid";
  const checkNameAvailabilityInputs = {
    name: "serviceName",
    type: "Microsoft.HealthcareApis/services",
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.services.checkNameAvailability(checkNameAvailabilityInputs);
  console.log(result);
}

checkNameAvailability().catch(console.error);
