const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a Linked Subscription resource.
 *
 * @summary Patch a Linked Subscription resource.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/LinkedSubscription/Patch.json
 */
async function patchALinkedSubscriptionResource() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const linkedSubscriptionName = "testLinkedSubscription";
  const resource = {
    linkedSubscriptionId: "104fbb77-2b0e-476a-83de-65ad8acd1f0b",
    location: "eastus",
    registrationResourceId:
      "/subscriptions/dd8597b4-8739-4467-8b10-f8679f62bfbf/resourceGroups/azurestack/providers/Microsoft.AzureStack/registrations/testRegistration",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.linkedSubscriptions.update(
    resourceGroup,
    linkedSubscriptionName,
    resource
  );
  console.log(result);
}

patchALinkedSubscriptionResource().catch(console.error);
