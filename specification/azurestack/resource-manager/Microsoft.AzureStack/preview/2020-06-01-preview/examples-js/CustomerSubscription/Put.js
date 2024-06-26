const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new customer subscription under a registration.
 *
 * @summary Creates a new customer subscription under a registration.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/CustomerSubscription/Put.json
 */
async function createsANewCustomerSubscriptionUnderARegistration() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const registrationName = "testregistration";
  const customerSubscriptionName = "E09A4E93-29A7-4EBA-A6D4-76202383F07F";
  const customerCreationParameters = {
    tenantId: "dbab3982-796f-4d03-9908-044c08aef8a2",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.customerSubscriptions.create(
    resourceGroup,
    registrationName,
    customerSubscriptionName,
    customerCreationParameters
  );
  console.log(result);
}

createsANewCustomerSubscriptionUnderARegistration().catch(console.error);
