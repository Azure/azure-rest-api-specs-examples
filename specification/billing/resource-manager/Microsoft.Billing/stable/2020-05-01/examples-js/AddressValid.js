const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates an address. Use the operation to validate an address before using it as soldTo or a billTo address.
 *
 * @summary Validates an address. Use the operation to validate an address before using it as soldTo or a billTo address.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AddressValid.json
 */
async function addressValid() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const address = {
    addressLine1: "1 Test Address",
    city: "bellevue",
    country: "us",
    postalCode: "12345",
    region: "wa",
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.address.validate(address);
  console.log(result);
}

addressValid().catch(console.error);
