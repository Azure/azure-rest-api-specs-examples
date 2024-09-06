const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates an address. Use the operation to validate an address before using it as soldTo or a billTo address.
 *
 * @summary Validates an address. Use the operation to validate an address before using it as soldTo or a billTo address.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/addressValidateInvalid.json
 */
async function addressValidateInvalid() {
  const parameters = {
    addressLine1: "1 Test",
    city: "bellevue",
    country: "us",
    postalCode: "12345",
    region: "wa",
  };
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential);
  const result = await client.address.validate(parameters);
  console.log(result);
}
