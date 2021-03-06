const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if a service instance name is available.
 *
 * @summary Check if a service instance name is available.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/CheckNameAvailabilityPost.json
 */
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
