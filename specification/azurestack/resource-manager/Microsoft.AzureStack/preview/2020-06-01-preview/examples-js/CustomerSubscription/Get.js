const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the specified product.
 *
 * @summary Returns the specified product.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/CustomerSubscription/Get.json
 */
async function returnsTheSpecifiedProduct() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const registrationName = "testregistration";
  const customerSubscriptionName = "E09A4E93-29A7-4EBA-A6D4-76202383F07F";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.customerSubscriptions.get(
    resourceGroup,
    registrationName,
    customerSubscriptionName
  );
  console.log(result);
}

returnsTheSpecifiedProduct().catch(console.error);
