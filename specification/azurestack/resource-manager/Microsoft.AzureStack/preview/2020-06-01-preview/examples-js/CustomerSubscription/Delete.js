const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a customer subscription under a registration.
 *
 * @summary Deletes a customer subscription under a registration.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/CustomerSubscription/Delete.json
 */
async function deletesACustomerSubscriptionUnderARegistration() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const registrationName = "testregistration";
  const customerSubscriptionName = "E09A4E93-29A7-4EBA-A6D4-76202383F07F";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.customerSubscriptions.delete(
    resourceGroup,
    registrationName,
    customerSubscriptionName
  );
  console.log(result);
}

deletesACustomerSubscriptionUnderARegistration().catch(console.error);
