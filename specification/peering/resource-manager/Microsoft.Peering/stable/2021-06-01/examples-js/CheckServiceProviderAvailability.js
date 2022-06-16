const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if the peering service provider is present within 1000 miles of customer's location
 *
 * @summary Checks if the peering service provider is present within 1000 miles of customer's location
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CheckServiceProviderAvailability.json
 */
async function checkIfPeeringServiceProviderIsAvailableInCustomerLocation() {
  const subscriptionId = "subId";
  const checkServiceProviderAvailabilityInput = {
    peeringServiceLocation: "peeringServiceLocation1",
    peeringServiceProvider: "peeringServiceProvider1",
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.checkServiceProviderAvailability(
    checkServiceProviderAvailabilityInput
  );
  console.log(result);
}

checkIfPeeringServiceProviderIsAvailableInCustomerLocation().catch(console.error);
