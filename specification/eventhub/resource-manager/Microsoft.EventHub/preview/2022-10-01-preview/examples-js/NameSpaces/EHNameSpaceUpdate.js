const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a namespace. Once created, this namespace's resource manifest is immutable. This operation is idempotent.
 *
 * @summary Creates or updates a namespace. Once created, this namespace's resource manifest is immutable. This operation is idempotent.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/EHNameSpaceUpdate.json
 */
async function namespacesUpdate() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "SampleSubscription";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ResurceGroupSample";
  const namespaceName = "NamespaceSample";
  const parameters = {
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/MicrosoftManagedIdentity/userAssignedIdentities/ud2":
          {},
      },
    },
    location: "East US",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.update(resourceGroupName, namespaceName, parameters);
  console.log(result);
}
